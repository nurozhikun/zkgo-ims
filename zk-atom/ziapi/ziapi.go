package ziapi

import (
	"reflect"

	"gitee.com/sienectagv/gozk/zlogger"
	"gitee.com/sienectagv/gozk/znet"
	"gitee.com/sienectagv/gozk/zproto"
	"gitee.com/sienectagv/gozk/zsql"
	"gitee.com/sienectagv/gozk/zutils"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi/apias"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi/apiauth"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi/apimap"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zinet/netirisbee"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf/protbee"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql"
)

type (
	ApiAuth = apiauth.ApiAuth
	ApiMap  = apimap.ApiMap
)

func NewApiAuth(db *zsql.DB) *ApiAuth {
	return &ApiAuth{DB: zisql.WrapupDbAuth(db)}
}
func NewApiMap(db *zsql.DB) *ApiMap {
	return &ApiMap{DB: zisql.WrapupDbMap(db)}
}

type IIrisBeeApi interface {
	// ProtBeeDb() zisql.IProtBeeDB
}

type IrisBeeApis struct {
	apis []IIrisBeeApi
}

func (iba *IrisBeeApis) InstallApi(api IIrisBeeApi) {
	iba.apis = append(iba.apis, api)
}

func (iba *IrisBeeApis) InstallBeeHandleByCmds(party znet.IrisParty, cmds ...int) {
	for _, cmd := range cmds {
		name, ok := zisql.BeeFuncNames[cmd]
		if !ok {
			continue
		}
		path, ok := apias.CmdPaths[cmd]
		if !ok {
			continue
		}
		for _, api := range iba.apis {
			fn := BeeHandleByFuncName(api, name)
			if nil == fn {
				continue
			}
			reqBody := apias.GetBeeReqBodyByCmd(cmd)
			ctxFn := func(ctx znet.IrisCtx) {
				var resMsg zproto.Message
				h, err := netirisbee.ParserHeader(ctx) //get header
				defer func() {
					netirisbee.SetHeader(ctx, h, err)
					if nil != resMsg {
						ctx.Text(zproto.MarshalString(resMsg)) //write to context
					}
				}()
				if nil != err {
					return
				}
				if nil != reqBody {
					if err = netirisbee.ParseBody(ctx, reqBody); nil != err {
						return
					}
				} else {
					reqBody = &zproto.EmptyMessage{}
				}
				resMsg, err = fn(h, reqBody) //get response data
			}
			party.Post(path, ctxFn)
			party.Get(path, ctxFn)
			return
		}
	}
}

func BeeHandleByFuncName(api interface{}, method string) func(h *protbee.Header, req zproto.Message) (zproto.Message, error) {
	if nil == api {
		return nil
	}
	value := reflect.ValueOf(api)
	if !value.IsValid() {
		return nil
	}
	zlogger.Info(value)
	vf := value.MethodByName(method)
	if !vf.IsValid() {
		return nil
	}
	// zlogger.Info(vf)
	return func(h *protbee.Header, req zproto.Message) (zproto.Message, error) {
		in := make([]reflect.Value, 2)
		in[0] = reflect.ValueOf(h)
		if nil != req {
			in[1] = reflect.ValueOf(req)
		} else {
			in[1] = reflect.New(reflect.TypeOf(req))
			// zlogger.Info(in[1].IsValid())
		}
		out := vf.Call(in)
		if len(out) < 2 {
			return nil, zutils.NewError(-1, "the response return is wrong")
		}
		msg, _ := out[0].Interface().(zproto.Message)
		err, _ := out[1].Interface().(error)
		return msg, err
	}
}
