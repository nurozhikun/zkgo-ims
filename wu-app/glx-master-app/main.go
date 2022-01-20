package main

import (
	"gitee.com/sienectagv/gozk/zsql"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi"

	"gitee.com/sienectagv/gozk/zutils"
)

func main() {
	waitGroup := zutils.NewLoopGroup()
	master := NewGlxMaster()
	//create and append db
	cfg := &zsql.Cfg{
		Type:     1, //Mysql
		Addr:     "127.0.0.1:3306",
		User:     "root",
		Password: "123456",
		Database: "ims_auth",
	}
	master.AppendDb("auth", cfg)
	//install wen handles
	party := master.app.Party("/auth")
	master.InstallBeeAdminApi(party, master.DB("auth"), ziapi.AuthCmds...)
	waitGroup.WaitForEnter("quit")
}
