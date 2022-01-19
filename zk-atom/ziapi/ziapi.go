package ziapi

import (
	"reflect"
	"zkgo-ims/zk-atom/ziapi/apias"

	"gitee.com/sienectagv/gozk/znet"
	"gitee.com/sienectagv/gozk/zproto"
	"gitee.com/sienectagv/gozk/zsql"
	"gitee.com/sienectagv/gozk/zutils"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi/apiauth"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi/apimap"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zinet/netirisbee"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf/protbee"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql"
)

type (
	Command           = apias.Command
	ApiAuth           = apiauth.ApiAuth
	ApiMap            = apimap.ApiMap
	BeeResponseMethod = func(h *protbee.Header, req zproto.Message) (zproto.Message, error)
)

func NewApiAuth(db *zsql.DB) *ApiAuth {
	return &ApiAuth{DB: zisql.WrapupDbAuth(db)}
}
func NewApiMap(db *zsql.DB) *ApiMap {
	return &ApiMap{DB: zisql.WrapupDbMap(db)}
}

type IIrisBeeApi interface {
	//the method of cmd must be func(h *protbee.Header, req zproto.Message) (zproto.Message, error)
	MethodNameOfCmd(cmd int) (string, bool)
	PathOfCmd(cmd int) (string, bool)
	ReqBodyOfCmd(cmd int) zproto.Message
}

type IrisBeeApis struct {
	apis []IIrisBeeApi
}

func (iba *IrisBeeApis) InstallApi(api IIrisBeeApi) {
	iba.apis = append(iba.apis, api)
}

func (iba *IrisBeeApis) InstallBeeHandlerOfCmds(party znet.IrisParty, cmds ...Command) {
	for _, api := range iba.apis {
		for _, cmd := range cmds {
			fn := BeeHandleByFuncName(api, cmd.MethodName)
			if nil == fn {
				continue
			}
			reqBody := cmd.BeeReqBody()
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
			party.Post(cmd.Path, ctxFn)
			party.Get(cmd.Path, ctxFn)
		}
	}
}

func (iba *IrisBeeApis) InstallBeeHandles(party znet.IrisParty, cmds ...int) {
	for _, api := range iba.apis {
		for _, cmd := range cmds {
			name, ok := api.MethodNameOfCmd(cmd)
			if !ok {
				continue
			}
			path, ok := api.PathOfCmd(cmd)
			if !ok {
				continue
			}
			fn := BeeHandleByFuncName(api, name)
			if nil == fn {
				continue
			}
			reqBody := api.ReqBodyOfCmd(cmd)
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
	vf := value.MethodByName(method)
	if !vf.IsValid() {
		return nil
	}
	var beeMethod BeeResponseMethod
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
