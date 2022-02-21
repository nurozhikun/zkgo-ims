package ziapi

import (
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf/protbee"

	"gitee.com/sienectagv/gozk/zproto"
)

const (
	CmdAuthLogin      = 1
	CmdAuthLogout     = 2
	CmdAuthTest       = 3
	CmdAuthGetUsers   = 4
	CmdAuthGetRoles   = 5
	CmdAuthAddUser    = 6
	CmdAuthDelUser    = 7
	CmdAuthUpdateUser = 8
	CmdMapThumbnails  = 1001
	CmdMapOneDetail   = 1002
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
	colAuthCommand(Command{
		Cmd:          CmdAuthTest,
		Path:         "/test",
		MethodName:   "BeeTest",
		FnBeeReqBody: func() zproto.Message { return &protbee.EmptyMessage{} },
	})
	colAuthCommand(Command{
		Cmd:          CmdAuthGetUsers,
		Path:         "/get_users",
		MethodName:   "BeeGetUsers",
		FnBeeReqBody: func() zproto.Message { return &protbee.UsersReq{} },
	})
	colAuthCommand(Command{
		Cmd:          CmdAuthGetRoles,
		Path:         "/get_roles",
		MethodName:   "BeeGetRoles",
		FnBeeReqBody: func() zproto.Message { return &protbee.RolesReq{} },
	})
	colAuthCommand(Command{
		Cmd:          CmdAuthAddUser,
		Path:         "/add_user",
		MethodName:   "BeeAddUser",
		FnBeeReqBody: func() zproto.Message { return &protbee.ManageUserReq{} },
	})
	colAuthCommand(Command{
		Cmd:          CmdAuthDelUser,
		Path:         "/del_user",
		MethodName:   "BeeDelUser",
		FnBeeReqBody: func() zproto.Message { return &protbee.ManageUserReq{} },
	})
	colAuthCommand(Command{
		Cmd:          CmdAuthUpdateUser,
		Path:         "/update_user",
		MethodName:   "BeeUpdateUser",
		FnBeeReqBody: func() zproto.Message { return &protbee.ManageUserReq{} },
	})
}

func colAuthCommand(c Command) {
	AuthCmds = append(AuthCmds, c)
	AllCmds[c.Cmd] = c
}
