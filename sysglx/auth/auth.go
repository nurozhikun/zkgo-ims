package auth

import (
	"gitee.com/sienectagv/gozk/zlogger"
	"gitee.com/sienectagv/gozk/zsql"
	"github.com/kataras/iris/v12"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql/zidbauth"
)

type Auth struct {
	App  *iris.Application
	db   *zidbauth.DB
	apis *ziapi.IrisBeeApis
}

func NewAuth(App *iris.Application) *Auth {
	return &Auth{
		App:  App,
		apis: &ziapi.IrisBeeApis{},
	}
}

func (a *Auth) InitDB(cfg *zsql.Cfg) *Auth {
	a.db = zidbauth.CreateDB(cfg)
	return a
}

func (a *Auth) InitApi(cmds ...int) *Auth {
	if a.App == nil {
		a.App = iris.New()
	}
	if a.db == nil {
		zlogger.Error("db is nil")
		return a
	}
	a.apis.InstallApi(ziapi.NewApiAuth(a.db.DB))
	a.InstallPath(cmds...)
	return a
}

func (a *Auth) InstallPath(cmds ...int) {
	// authParty := a.App.Party("/auth")
	// {
	// 	a.apis.InstallBeeHandles(authParty, cmds...)
	// }
}
