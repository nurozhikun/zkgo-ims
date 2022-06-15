package main

import (
	"fmt"

	"gitee.com/sienectagv/gozk/zlogger"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi/ziclt"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf"
)

type TsBeeReqs struct {
	*ziclt.ZiHttpBee
}

func NewTsBeeReqs() *TsBeeReqs {
	return &TsBeeReqs{ZiHttpBee: ziclt.NewHttpBee()}
}

func (bee *TsBeeReqs) RunAuth() {
	// fmt.Printf("bee: %v\n", *bee)
	cmds := ziclt.MakeAuthReqCommands()
	fmt.Println(cmds)
	for _, cmd := range cmds {
		r, err := zipbf.BeeHttpPost(bee.HttpClient, bee.BeeHeader, &cmd)
		zlogger.Info(r.Status(), err, r.Result())
	}
}
