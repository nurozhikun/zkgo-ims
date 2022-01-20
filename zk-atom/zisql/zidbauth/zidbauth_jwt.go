package zidbauth

import (
	"gitee.com/sienectagv/gozk/zjwt"
)

///为了简化权限管理，权限只管理一级目录，比如 path /auth/login; /task/add; 只控制 auth, task 这些一级目录

var SecretKey = []byte("Sienect@2021$jwt")

func TokenOfRoles(user *User) (string, error) {
	roles := make([]uint, 0, len(user.Roles))
	for _, r := range user.Roles {
		roles = append(roles, r.Model.ID)
	}
	return zjwt.TokenOfRoles(user.Name, SecretKey, roles)
}

func ParseTokenOfRoles(tokenStr string) (user string, roles []uint, err error) {
	return zjwt.ParseTokenOfRoles(SecretKey, tokenStr)
}
