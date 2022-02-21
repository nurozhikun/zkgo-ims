package zidbauth

import (
	"errors"
	"strings"

	"github.com/kataras/iris/v12/context"

	"gitee.com/sienectagv/gozk/zjwt"
	"gitee.com/sienectagv/gozk/zlogger"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf/protbee"
)

var jwtDB *DB

func SetJwtDB(db *DB) {
	jwtDB = db
}

func JwtAuthCheck(header *protbee.Header, ctx context.Context) bool {
	username, perms, err := ParseTokenOfPermissions(header.Jwt)
	if err != nil {
		zlogger.Error(err)
		return false
	}
	for _, v := range perms {
		if strings.HasPrefix(ctx.GetCurrentRoute().Path(), v) {
			return true
		}
	}

	// 未通过则在数据库拉取
	if jwtDB == nil {
		zlogger.Error(errors.New("JWT DB is nil"))
		return false
	}
	user := &User{}
	if err = jwtDB.Where("name = ?", username).First(user).Error; err != nil {
		zlogger.Error(err)
		return false
	}
	if err := jwtDB.Model(user).Association("Roles").Find(&(user.Roles)); err != nil {
		zlogger.Error(err)
		return false
	}
	// 填充permissions
	for _, v := range user.Roles {
		if err := jwtDB.Model(v).Association("Permissions").Find(&(v.Permissions)); err != nil {
			zlogger.Error(err)
			return false
		}
	}
	newPerms := make([]string, 0)
	for _, role := range user.Roles {
		for _, perm := range role.Permissions {
			newPerms = append(newPerms, perm.Url)
		}
	}
	newJwt, err := zjwt.TokenOfPermissons(username, SecretKey, newPerms)
	if err != nil {
		zlogger.Error(err)
		return false
	}
	header.Jwt = newJwt
	for _, v := range newPerms {
		if strings.HasPrefix(ctx.GetCurrentRoute().Path(), v) {
			return true
		}
	}

	return false
}
