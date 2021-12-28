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
	return &ApiMap{dbMap: zisql.WrapupDbMap(db)}
}

type ApiMap struct {
	dbMap *zidbmap.DB
}

func (m *ApiMap) IrisBeeHandle(key int) func(ctx znet.IrisCtx) {
	switch key {
	case apias.Cmd_MapThumbnails:
		return getIrisHandle(m, (*ApiMap).resBeeThumbnails)
	default:
		return nil
	}
}

func getIrisHandle(m *ApiMap, fn func(a *ApiMap, header *zipbf.BeeHeader) (zproto.Message, error)) func(ctx znet.IrisCtx) {
	return func(ctx znet.IrisCtx) {
		h, err := netirisbee.ParserHeader(ctx)
		if nil != err {
			netirisbee.SetHeader(ctx, nil)
			netirisbee.ResponseError(ctx, err, netas.ResCode_UnknownError)
		} else {
			netirisbee.SetHeader(ctx, h)
			resMsg, err := fn(m, h)
			if nil != err || nil == resMsg {
				netirisbee.ResponseError(ctx, err, netas.ResCode_UnknownError)
			} else {
				ctx.Text(zproto.MarshalString(resMsg))
			}
		}
	}
}

func (m *ApiMap) resBeeThumbnails(header *zipbf.BeeHeader) (zproto.Message, error) {
	body := &protbee.ResMapThumbnails{}
	//TODO read data from zidbmap then fill body
	return body, nil
}

func (m *ApiMap) resBeeOneDetail(header *zipbf.BeeHeader) (zproto.Message, error) {
	return nil, nil
}
