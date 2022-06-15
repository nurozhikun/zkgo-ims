package xincapi

import (
	"github.com/kataras/iris/v12"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapp"
)

type XincApi struct {
	ziapp.ZiAppBee
}

func NewXincApi() *XincApi {
	app := &XincApi{}
	app.Init(iris.New())
	// app.InstallBeeProtAuthDefault() //api of default auth
	return app
}
