package main

import (
	"gitee.com/sienectagv/gozk/zsql"
	"gitee.com/sienectagv/gozk/zutils"
	"github.com/nurozhikun/zkgo-ims/zk-planet/zkamr"
)

func main() {
	waitGroup := zutils.NewLoopGroup()
	amr := &zkamr.ZkAmr{}
	defer amr.FreeAll()
	// init db
	cfg := &zsql.Cfg{
		Type:     zsql.TypeSqlite,
		Database: "user/map.db",
	}
	amr.InitDB(cfg)
	amr.InitApi()
	amr.StartDevices()
	waitGroup.WaitForEnter("quit")
}
