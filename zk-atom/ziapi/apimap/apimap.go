package apimap

import (
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi/apias"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql/zidbmap"
)

type ApiMap struct {
	apias.ApiBase
	*zidbmap.DB //需要用继承
}

func (a *ApiMap) ProtBeeDb() zisql.IProtBeeDB {
	return a.DB
}

/*
func (m *ApiMap) IrisBeeHandle(cmd int) func(ctx znet.IrisCtx) {
	switch cmd {
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
		}
		resMsg, err = fn(m, h, reqBody) //get response data
	}
}
*/
