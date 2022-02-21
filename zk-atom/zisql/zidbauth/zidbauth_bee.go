package zidbauth

import (
	"errors"

	"gitee.com/sienectagv/gozk/zlogger"
	"gitee.com/sienectagv/gozk/zproto"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf/protbee"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql/zidb"
)

// 暂未使用
func (db *DB) ProtBeeResponse(cmd int, h *protbee.Header, res zproto.Message) {
}

// 写登录逻辑(和之前的唯一不同就是入参和返回值都使用message接口)
// 然后参照之前的程序 集成到galaxy中
// header传入指针，则只需要返回res即可
func (db *DB) BeeAuthLogin(header *zipbf.BeeHeader, reqBody zproto.Message) (message zproto.Message, err error) {
	defer func() {
		if err != nil {
			zlogger.Error(err)
		}
	}()

	req, ok := reqBody.(*protbee.UserReq)
	if !ok {
		return nil, errors.New("input type error, need UserReq")
	}

	if req.User == "" {
		return nil, errors.New("user cann't be null")
	}

	user, err := db.CheckInUser(req.User, req.Password)
	if err != nil {
		zlogger.Error(err)
		return nil, err
	}
	// 把权限写入token, 此前是写入Role
	szToken, err := TokenOfPermissons(user)
	if err != nil {
		return nil, err
	}

	// 暂只填充jwt与username
	res := &protbee.UserRes{
		Jwt:  szToken,
		User: user.Name,
	}
	header.Jwt = szToken
	return res, nil
}

func (db *DB) BeeTest(header *zipbf.BeeHeader, reqBody zproto.Message) (message zproto.Message, err error) {
	return &protbee.Test{
		Test: "123",
	}, nil
}

func (db *DB) BeeGetUsers(header *zipbf.BeeHeader, reqBody zproto.Message) (message zproto.Message, err error) {
	defer func() {
		if err != nil {
			zlogger.Error(err)
		}
	}()

	users := []User{}
	// 嵌套预加载
	if err := db.Preload("Roles.Permissions").Preload("Roles").Find(&users).Error; err != nil {
		return nil, err
	}
	usersRes := protbee.UsersRes{}
	for _, v := range users {
		usersRes.Users = append(usersRes.Users, &protbee.UserRes{
			User: v.Name,
			Jwt: func() string {
				szToken, _ := TokenOfPermissons(&v)
				return szToken
			}(),
			Tel: v.Mobile,
		})
	}

	return &usersRes, nil
}

func (db *DB) BeeGetRoles(header *zipbf.BeeHeader, reqBody zproto.Message) (message zproto.Message, err error) {
	defer func() {
		if err != nil {
			zlogger.Error(err)
		}
	}()

	roles := []Role{}
	if err := db.Find(&roles).Error; err != nil {
		return nil, err
	}
	protRoles := protbee.RolesRes{}
	for _, v := range roles {
		protRoles = append(protRoles.Roles, &protbee.Role{
			Id:   int64(v.Model.ID),
			Code: v.Code,
			Name: v.Name,
		})
	}

	return &protRoles, nil
}

func (db *DB) BeeAddUser(header *zipbf.BeeHeader, reqBody zproto.Message) (message zproto.Message, err error) {
	defer func() {
		if err != nil {
			zlogger.Error(err)
		}
	}()

	userReq, ok := reqBody.(*protbee.ManageUserReq)
	if !ok {
		return nil, errors.New("input type error, need ManageUserReq")
	}
	if userReq.User == "" {
		return nil, return errors.New("name is nil")
	}
	if userReq.Password == "" {
		userReq.Password = "sienect"
	}
	sum, salt, _ := zidb.SaltPassCreate(userReq.User, userReq.Password)
	user := User{
		Name: userReq.User,
		Password: sum,
		Mobile: userReq.Mobile,
		Salt: salt,
	}
	roles := []*Role{}
	for _, v := range reqBody.User.Roles {
		roles = append(roles, &Role{
			Model: gorm.Model{
				ID: unit(v.ID),
			},
			Name: v.Name,
			Code: v.Code,
		})
	}
	user.Roles = roles
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &protbee.ManageUserRes{}, nil
}

func (db *DB) BeeDelUser(header *zipbf.BeeHeader, reqBody zproto.Message) (message zproto.Message, err error) {
	defer func() {
		if err != nil {
			zlogger.Error(err)
		}
	}()

	userReq, ok := reqBody.(*protbee.ManageUserReq)
	if !ok {
		return nil, errors.New("input type error, need ManageUserReq")
	}
	if err := db.Where("name = ?", userReq.User).Delete(&User{}); err != nil {
		return nil, err
	}
	return &protbee.ManageUserRes{}, nil
}

func (db *DB) BeeUpdateUser(header *zipbf.BeeHeader, reqBody zproto.Message) (message zproto.Message, err error) {
	defer func() {
		if err != nil {
			zlogger.Error(err)
		}
	}()

	userReq, ok := reqBody.(*protbee.ManageUserReq)
	if !ok {
		return nil, errors.New("input type error, need ManageUserReq")
	}

	user := &User{}
	if err := db.Where("name = ?", userReq.User).First(&user).Error; err != nil {
		return nil, err
	}
	user.Mobile = userReq.Mobile
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Updates(&user).Error; err != nil {
			return err
		}
		roles := []*Role{}
		for _, v := range userReq.Roles {
			roles = append(roles, &Role{
				Model: gorm.Model{
					ID: uint(v.Id),
				},
				Name: v.Name,
				Code: v.Code,
			})
		}
		// 替换关联
		if err := tx.Model(&user).Association("Roles").Replace(roles); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &protbee.ManageUserRes{}, nil
}
