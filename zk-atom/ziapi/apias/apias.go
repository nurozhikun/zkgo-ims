package apias

import (
	"gitee.com/sienectagv/gozk/zproto"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf/protbee"
)

const (
	// for apiauth
	Cmd_AuthLogin  = 1
	Cmd_AuthLogout = 2
	// for apimaps
	Cmd_MapThumbnails = 1001
	Cmd_MapOneDetail  = 1002
)

var CmdPaths = map[int]string{
	Cmd_AuthLogin:     "/login",
	Cmd_AuthLogout:    "/logout",
	Cmd_MapThumbnails: "/thumbnails",
	Cmd_MapOneDetail:  "/onedetail",
}

var BeeFuncNames = map[int]string{
	Cmd_AuthLogin:     "BeeAuthLogin",
	Cmd_AuthLogout:    "BeeAuthLogout",
	Cmd_MapThumbnails: "BeeMapThumbnails",
	Cmd_MapOneDetail:  "BeeMapOneDetail",
}

type ApiBase struct{}

func (a *ApiBase) ReqBodyOfCmd(cmd int) zproto.Message {
	switch cmd {
	case Cmd_AuthLogin:
		return &protbee.UserReq{}
	case Cmd_AuthLogout:
		// return &protbee.Logout{}
		return nil
	case Cmd_MapThumbnails:
		return nil
	case Cmd_MapOneDetail: //
		return &protbee.Thumbnail{}
	default:
		return nil
	}
}

func (a *ApiBase) MethodNameOfCmd(cmd int) (string, bool) {
	s, ok := BeeFuncNames[cmd]
	return s, ok
}

func (a *ApiBase) PathOfCmd(cmd int) (string, bool) {
	s, ok := CmdPaths[cmd]
	return s, ok
}
