package ziapi

import (
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf/protbee"

	"gitee.com/sienectagv/gozk/zproto"
)

const (
	CmdAuthLogin     = 1
	CmdAuthLogout    = 2
	CmdMapThumbnails = 1001
	CmdMapOneDetail  = 1002
)

var (
	AllCmds  = map[int]Command{}
	AuthCmds = []Command{}
	MapCmds  = []Command{}
)

func init() {
	//auth command
	colAuthCommand(Command{
		Cmd:          CmdAuthLogin,
		Path:         "/auth/login",
		MethodName:   "BeeAuthLogin",
		FnBeeReqBody: func() zproto.Message { return &protbee.UserRes{} },
	})
	colAuthCommand(Command{
		Cmd:          CmdAuthLogout,
		Path:         "/auth/logout",
		MethodName:   "BeeAuthLogout",
		FnBeeReqBody: func() zproto.Message { return &protbee.UserRes{} },
	})
}

func colAuthCommand(c Command) {
	AuthCmds = append(AuthCmds, c)
	AllCmds[c.Cmd] = c
}
