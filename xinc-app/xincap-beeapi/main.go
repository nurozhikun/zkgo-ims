package main

import (
	"context"
	"fmt"

	"gitee.com/sienectagv/gozk/zsql"
	"gitee.com/sienectagv/gozk/zutils"
	"github.com/nurozhikun/zkgo-ims/xinc-app/xincap-beeapi/xincapi"
)

func main() {
	waitGroup := zutils.NewLoopGroup()
	app := xincapi.NewXincApi()
	cfg := &zsql.Cfg{
		Type:     zsql.TypeMysqlite,
		Addr:     "127.0.0.1:3306",
		User:     "root",
		Password: "123456",
		Database: "xinc_auth",
	}
	// fmt.Println(*cfg)
	// app.AppendDb(ziapp.DbKeyAuth, cfg)
	err := app.InstallBeeProtAuthDefault(cfg)
	fmt.Println(err)
	waitGroup.AddAsyncBlock(
		func() {
			app.Listen("0.0.0.0:8000")
		},
		func() {
			app.Shutdown(context.TODO())
		})
	waitGroup.WaitForEnter("quit")
}
