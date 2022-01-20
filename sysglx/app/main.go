package main

import (
	"context"

	"gitee.com/sienectagv/gozk/zsql"
	"gitee.com/sienectagv/gozk/zutils"
	"github.com/kataras/iris/v12"
	"github.com/nurozhikun/zkgo-ims/sysglx/auth"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi/apias"
)

func main() {
	waitGroup := zutils.NewLoopGroup()
	app := iris.New()
	au := auth.NewAuth(app)
	cfg := &zsql.Cfg{
		Type:     zsql.TypeMysqlite,
		Addr:     "127.0.0.1:3306",
		User:     "root",
		Password: "mysql",
		Database: "galaxy_auth",
	}
	au.InitDB(cfg)
	au.InitApi(apias.Cmd_AuthLogin)
	waitGroup.AddAsyncBlock(
		func() {
			app.Listen("0.0.0.0:8000")
		},
		func() {
			app.Shutdown(context.TODO())
		},
	)
	waitGroup.WaitForEnter("quit")
}
