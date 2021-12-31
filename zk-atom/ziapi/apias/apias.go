package apias

import (
	"gitee.com/sienectagv/gozk/zproto"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf/protbee"
)

const (
	//for apiauth
	Cmd_AuthLogin  = 1
	Cmd_AuthLogout = 2
	//for apimaps
	Cmd_MapThumbnails = 1001
	Cmd_MapOneDetail  = 1002
)

var CmdPaths = map[int]string{
	Cmd_AuthLogin:     "/login",
	Cmd_AuthLogout:    "/logout",
	Cmd_MapThumbnails: "/thumbnails",
	Cmd_MapOneDetail:  "/onedetail",
}

func GetBeeReqBodyByCmd(cmd int) zproto.Message {
	switch cmd {
	case Cmd_AuthLogin:
		return &protbee.Login{}
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

// func GetIrisHandle(
// 	reqBody zproto.Message,
// 	fn func(h *zipbf.BeeHeader, reqBody zproto.Message) (zproto.Message, error),
// ) func(ctx znet.IrisCtx) {
// 	return func(ctx znet.IrisCtx) {
// 		var resMsg zproto.Message
// 		h, err := netirisbee.ParserHeader(ctx) //get header
// 		defer func() {
// 			netirisbee.SetHeader(ctx, h, err)
// 			if nil != resMsg {
// 				ctx.Text(zproto.MarshalString(resMsg)) //write to context
// 			}
// 		}()
// 		if nil != err {
// 			return
// 		}
// 		if nil != reqBody {
// 			if err = netirisbee.ParseBody(ctx, reqBody); nil != err {
// 				return
// 			}
// 		}
// 		resMsg, err = fn(h, reqBody) //get response data
// 	}
// }
