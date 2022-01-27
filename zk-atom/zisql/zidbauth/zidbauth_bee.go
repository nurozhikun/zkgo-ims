package zidbauth

import (
	"errors"

	"gitee.com/sienectagv/gozk/zlogger"
	"gitee.com/sienectagv/gozk/zproto"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf/protbee"
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

	// TODO 还需要检查reqBody怎么使用

	req, ok := reqBody.(*protbee.UserReq)
	if !ok {
		return nil, errors.New("input type error, need UserReq")
	}

	if req.User == "" {
		return nil, errors.New("user cann't be null")
	}

	user, err := db.CheckInUser(req.User, req.Password)
	if err != nil {
		return nil, err
	}

	szToken, err := TokenOfRoles(user)
	if err != nil {
		return nil, err
	}

	// 暂只填充jwt与username
	res := &protbee.UserRes{
		Jwt:  szToken,
		User: user.Name,
	}
	return res, nil
}
