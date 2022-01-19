package main

import (
	"fmt"

	"gitee.com/sienectagv/gozk/zsql"

	"gitee.com/sienectagv/gozk/zutils"
)

func main() {
	waitGroup := zutils.NewLoopGroup()

	cfg := &zsql.Cfg{
		Type:     1, //Mysql
		Addr:     "127.0.0.1:3306",
		User:     "root",
		Password: "123456",
		Database: "galaxy",
	}
	fmt.Println(cfg)
	waitGroup.WaitForEnter("quit")
}
