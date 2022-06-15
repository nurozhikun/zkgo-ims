package ziapi

import (
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf/protbee"

	"gitee.com/sienectagv/gozk/znet"
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

// type ReqBodyType = interface{}

// type Command struct {
// 	Cmd        int
// 	Path       string
// 	MethodName string
// 	// 要用函数的原因是每次都要生成一个对象，不能用共享的一个对象
// 	FnRequestBody func() ReqBodyType
// }

type Command = znet.Command

// func (c *Command) BeeProtoReqBody() zproto.Message {
// 	if nil == c.FnRequestBody {
// 		return nil
// 	}
// 	r := c.FnRequestBody()
// 	if nil == r {
// 		return nil
// 	}
// 	req, _ := r.(zproto.Message)
// 	return req
// }

var (
	AllBeeCmds  = map[int]Command{}
	AuthBeeCmds = []Command{}
	MapBeeCmds  = []Command{}
)

func CommandOf(cmd int) Command {
	c, ok := AllBeeCmds[cmd]
	if !ok {
		return Command{Cmd: cmd}
	} else {
		return c
	}
}

func init() {
	// auth command
	colAuthCommand(Command{
		Cmd:        CmdAuthLogin,
		Path:       "/login",
		MethodName: "BeeAuthLogin",
		// 这里出错，应该使用UserReq{}
		FnRequestBody: func() znet.ReqBodyType { return &protbee.UserReq{} },
	})
	colAuthCommand(Command{
		Cmd:           CmdAuthLogout,
		Path:          "/logout",
		MethodName:    "BeeAuthLogout",
		FnRequestBody: func() znet.ReqBodyType { return &protbee.UserRes{} },
	})
	colAuthCommand(Command{
		Cmd:           CmdAuthTest,
		Path:          "/test",
		MethodName:    "BeeTest",
		FnRequestBody: func() znet.ReqBodyType { return &protbee.EmptyMessage{} },
	})
	colAuthCommand(Command{
		Cmd:           CmdAuthGetUsers,
		Path:          "/get_users",
		MethodName:    "BeeGetUsers",
		FnRequestBody: func() znet.ReqBodyType { return &protbee.UsersReq{} },
	})
	colAuthCommand(Command{
		Cmd:           CmdAuthGetRoles,
		Path:          "/get_roles",
		MethodName:    "BeeGetRoles",
		FnRequestBody: func() znet.ReqBodyType { return &protbee.RolesReq{} },
	})
	colAuthCommand(Command{
		Cmd:           CmdAuthAddUser,
		Path:          "/add_user",
		MethodName:    "BeeAddUser",
		FnRequestBody: func() znet.ReqBodyType { return &protbee.ManageUserReq{} },
	})
	colAuthCommand(Command{
		Cmd:           CmdAuthDelUser,
		Path:          "/del_user",
		MethodName:    "BeeDelUser",
		FnRequestBody: func() znet.ReqBodyType { return &protbee.ManageUserReq{} },
	})
	colAuthCommand(Command{
		Cmd:           CmdAuthUpdateUser,
		Path:          "/update_user",
		MethodName:    "BeeUpdateUser",
		FnRequestBody: func() znet.ReqBodyType { return &protbee.ManageUserReq{} },
	})
}

func colAuthCommand(c Command) {
	AuthBeeCmds = append(AuthBeeCmds, c)
	AllBeeCmds[c.Cmd] = c
}
