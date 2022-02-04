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
	// auth command
	colAuthCommand(Command{
		Cmd:        CmdAuthLogin,
		Path:       "/login",
		MethodName: "BeeAuthLogin",
		// 这里出错，应该使用UserReq{}
		FnBeeReqBody: func() zproto.Message { return &protbee.UserReq{} },
	})
	colAuthCommand(Command{
		Cmd:          CmdAuthLogout,
		Path:         "/logout",
		MethodName:   "BeeAuthLogout",
		FnBeeReqBody: func() zproto.Message { return &protbee.UserRes{} },
	})
}

func colAuthCommand(c Command) {
	AuthCmds = append(AuthCmds, c)
	AllCmds[c.Cmd] = c
}
