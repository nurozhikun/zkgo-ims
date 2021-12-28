package ziapi

import (
	"gitee.com/sienectagv/gozk/znet"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi/apias"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi/apimap"
)

type (
	ApiMap = apimap.ApiMap
)

var (
	NewApiMap = apimap.NewApiMap
)

type IIrisBeeApi interface {
	IrisBeeHandle(key int) func(ctx znet.IrisCtx)
}

func InstallIrisBeeHandles(api IIrisBeeApi, party znet.IrisParty, cmds ...int) {
	for _, cmd := range cmds {
		installOneIrisBeeHandle(api, party, cmd)
	}
}

func installOneIrisBeeHandle(api IIrisBeeApi, party znet.IrisParty, cmd int) {
	f := api.IrisBeeHandle(cmd)
	if nil == f {
		return
	}
	s, ok := apias.CmdPaths[cmd]
	if !ok {
		return
	}
	party.Post(s, f)
	party.Get(s, f)
}
