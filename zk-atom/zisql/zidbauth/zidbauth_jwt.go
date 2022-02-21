package zidbauth

import (
	"gitee.com/sienectagv/gozk/zjwt"
)

///为了简化权限管理，权限只管理一级目录，比如 path /auth/login; /task/add; 只控制 auth, task 这些一级目录

var SecretKey = []byte("Sienect@2021$jwt")

func TokenOfPermissons(user *User) (string, error) {
	perms := make([]string, 0)
	for _, role := range user.Roles {
		for _, perm := range role.Permissions {
			perms = append(perms, perm.Url)
		}
	}
	return zjwt.TokenOfPermissons(user.Name, SecretKey, perms)
}

func ParseTokenOfPermissions(tokenStr string) (user string, perms []string, err error) {
	return zjwt.ParseTokenOfPermissions(SecretKey, tokenStr)
}

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
