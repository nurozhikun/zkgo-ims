package apimap

import (
	"gitee.com/sienectagv/gozk/znet"
	"gitee.com/sienectagv/gozk/zproto"
	"gitee.com/sienectagv/gozk/zsql"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi/apias"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zinet/netas"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zinet/netirisbee"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf/protbee"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql/zidbmap"
)

func NewApiMap(db *zsql.DB) *ApiMap {
	return &ApiMap{DB: zisql.WrapupDbMap(db)}
}

type ApiMap struct {
	*zidbmap.DB //需要用继承
}

func (m *ApiMap) IrisBeeHandle(key int) func(ctx znet.IrisCtx) {
	switch key {
	case apias.Cmd_MapThumbnails:
		return getIrisHandle(m, nil, (*ApiMap).BeeResThumbnails)
	case apias.Cmd_MapOneDetail:
		return getIrisHandle(m, &protbee.Thumbnail{}, (*ApiMap).BeeResOneDetail)
	default:
		return nil
	}
}

func getIrisHandle(m *ApiMap,
	reqBody zproto.Message,
	fn func(a *ApiMap, header *zipbf.BeeHeader, reqBody zproto.Message) (zproto.Message, error)) func(ctx znet.IrisCtx) {
	return func(ctx znet.IrisCtx) {
		h, err := netirisbee.ParserHeader(ctx) //get header
		if nil != err {
			netirisbee.SetHeader(ctx, nil)
			netirisbee.ResponseError(ctx, err, netas.ResCode_UnknownError)
			return
		}
		netirisbee.SetHeader(ctx, h)
		if nil != reqBody {
			err = netirisbee.ParseBody(ctx, reqBody) //get body
			if nil != err {
				netirisbee.ResponseError(ctx, err, netas.ResCode_UnknownError)
				return
			}
		}
		resMsg, err := fn(m, h, reqBody) //get response data
		if nil != err || nil == resMsg {
			netirisbee.ResponseError(ctx, err, netas.ResCode_UnknownError)
		} else {
			ctx.Text(zproto.MarshalString(resMsg)) //write to context
		}
	}
}

// func (m *ApiMap) BeeResThumbnails(header *zipbf.BeeHeader, reqBody zproto.Message) (zproto.Message, error) {
// 	body := &protbee.ResMapThumbnails{}
// 	//TODO read data from zidbmap then fill body
// 	return body, nil
// }
