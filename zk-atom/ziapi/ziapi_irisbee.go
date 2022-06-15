package ziapi

import (
	"errors"
	"reflect"

	"gitee.com/sienectagv/gozk/zlogger"
	"gitee.com/sienectagv/gozk/znet"
	"gitee.com/sienectagv/gozk/zproto"
	"gitee.com/sienectagv/gozk/zutils"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zinet/netirisbee"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf/protbee"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql/zidbauth"
)

func (iba *HttpHandles) InstallBeeProtobufOfCmds(party znet.IrisParty, cmds ...Command) {
	for _, api := range iba.handles {
		for _, cmd := range cmds {
			fn := BeeHandleByFuncName(api, cmd.MethodName)
			if nil == fn {
				continue
			}
			// reqBody := cmd.BeeProtoReqBody()
			reqBody := zipbf.BeeReqBodyFromCmd(&cmd)
			Cmd := cmd.Cmd
			ctxFn := func(ctx znet.IrisCtx) {
				var resMsg zproto.Message
				h, err := netirisbee.ParserHeader(ctx) // get header
				if nil != err {
					zlogger.Error(err)
					return
				}

				defer func() {
					netirisbee.SetHeader(ctx, h, err)
					if nil != resMsg {
						// 写入字符串(实质为json)
						ctx.Text(zproto.MarshalString(resMsg)) // write to context
					}
					// 判断返回值是不是nil
					if nil == resMsg {
						ctx.Text("nil return")
					}
					if err != nil {
						h.Error = err.Error()
						h.Code = -1
					}
				}()

				// 根据接口CMD 做authCheck
				if Cmd != CmdAuthLogin {
					// 除登录外都需要authCheck
					// h.Jwt AuthCheck(逻辑未 通过则放行 未通过从数据库拉取确认一遍)
					if !zidbauth.JwtAuthCheck(h, ctx) {
						err = errors.New("the path is not permitted")
						zlogger.Error(err)
						return
					}
				}

				if nil != reqBody {
					// 请求参数一定要准确
					if err = netirisbee.ParseBody(ctx, reqBody); nil != err {
						zlogger.Error(err)
						return
					}
				} else {
					reqBody = &protbee.EmptyMessage{}
				}
				resMsg, err = fn(h, reqBody) // get response data
				ctx.Next()
			}
			party.Post(cmd.Path, ctxFn)
			party.Get(cmd.Path, ctxFn)
		}
	}
}

// func (iba *HttpHandles) InstallBeeHandles(party znet.IrisParty, cmds ...int) {
// 	for _, api := range iba.handles {
// 		for _, cmd := range cmds {
// 			name, ok := api.MethodNameOfCmd(cmd)
// 			if !ok {
// 				continue
// 			}
// 			path, ok := api.PathOfCmd(cmd)
// 			if !ok {
// 				continue
// 			}
// 			fn := BeeHandleByFuncName(api, name)
// 			if nil == fn {
// 				continue
// 			}
// 			reqBody := api.ReqBodyOfCmd(cmd)
// 			ctxFn := func(ctx znet.IrisCtx) {
// 				var resMsg zproto.Message
// 				h, err := netirisbee.ParserHeader(ctx) // get header
// 				defer func() {
// 					netirisbee.SetHeader(ctx, h, err)
// 					if nil != resMsg {
// 						ctx.Text(zproto.MarshalString(resMsg)) // write to context
// 					}
// 				}()
// 				if nil != err {
// 					return
// 				}
// 				if nil != reqBody {
// 					if err = netirisbee.ParseBody(ctx, reqBody); nil != err {
// 						return
// 					}
// 				} else {
// 					reqBody = &zproto.EmptyMessage{}
// 				}
// 				resMsg, err = fn(h, reqBody) // get response data
// 			}
// 			party.Post(path, ctxFn)
// 			party.Get(path, ctxFn)
// 		}
// 	}
// }

func BeeHandleByFuncName(api interface{}, method string) func(h *protbee.Header, req zproto.Message) (zproto.Message, error) {
	if nil == api {
		return nil
	}
	value := reflect.ValueOf(api)
	if !value.IsValid() {
		return nil
	}
	vf := value.MethodByName(method)
	if !vf.IsValid() {
		return nil
	}
	var beeMethod beeResponseMethod
	if vf.Type() != reflect.TypeOf(beeMethod) {
		return nil
	}
	return func(h *protbee.Header, req zproto.Message) (zproto.Message, error) {
		in := make([]reflect.Value, 2)
		in[0] = reflect.ValueOf(h)
		in[1] = reflect.ValueOf(req)
		out := vf.Call(in)
		if len(out) < 2 {
			return nil, zutils.NewError(-1, "the response return is wrong")
		}
		msg, _ := out[0].Interface().(zproto.Message)
		err, _ := out[1].Interface().(error)
		return msg, err
	}
}
