package main

import (
	"gitee.com/sienectagv/gozk/zutils"
	"github.com/nurozhikun/zkgo-ims/zk-planet/zipamr"
)

func main() {
	waitGroup := zutils.NewLoopGroup()
	amr := &zipamr.ZipAmr{}
	//init db
	// cfg := &zsql.Cfg{
	// 	Type:     zsql.TypeSqlite,
	// 	Database: "user/map.db",
	// }
	// amr.InitDB(cfg).
	amr.InitDevs()
	waitGroup.WaitForEnter("quit")
	amr.FreeAll()
}
