package ziclt

import (
	"gitee.com/sienectagv/gozk/znet"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf/protbee"
)

func MakeAuthReqCommands() []ziapi.Command {
	cmds := []ziapi.Command{}
	c := ziapi.CommandOf(ziapi.CmdAuthLogin)
	c.FnRequestBody = func() znet.ReqBodyType {
		return &protbee.UserReq{
			User:     "wuzhikun",
			Password: "Wzk123456",
		}
	}
	cmds = append(cmds, c)
	return cmds
}
