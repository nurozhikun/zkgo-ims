package main

import (
	"gitee.com/sienectagv/gozk/zsql"

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
	waitGroup.WaitForEnter("quit")
}
