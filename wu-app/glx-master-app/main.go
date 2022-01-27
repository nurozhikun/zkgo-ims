package main

import (
	"context"

	"gitee.com/sienectagv/gozk/zsql"
	"github.com/nurozhikun/zkgo-ims/wu-app/glx-master-app/master"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi"

	"gitee.com/sienectagv/gozk/zutils"
)

func main() {
	waitGroup := zutils.NewLoopGroup()
	glxMaster := master.NewGlxMaster()
	// create and append db
	cfg := &zsql.Cfg{
		Type:     zsql.TypeMysqlite,
		Addr:     "127.0.0.1:3306",
		User:     "root",
		Password: "mysql",
		Database: "galaxy_auth",
	}
	// append zsql db
	glxMaster.AppendDb("auth", cfg)
	// install wen handles
	party := glxMaster.App.Party("/auth")
	// 把所有步骤统一起来
	glxMaster.InstallBeeAdminApi(party, glxMaster.DB("auth"), ziapi.AuthCmds...)
	// 起asyncBlock listen
	waitGroup.AddAsyncBlock(
		func() {
			glxMaster.App.Listen("0.0.0.0:8000")
		},
		func() {
			glxMaster.App.Shutdown(context.TODO())
		},
	)
	waitGroup.WaitForEnter("quit")
}
