package apiauth

import (
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql/zidbauth"
)

type ApiAuth struct {
	*zidbauth.DB
}

func (a *ApiAuth) ProtBeeDb() zisql.IProtBeeDB {
	return a.DB
}

// func (a *ApiAuth) IrisBeeHandle(cmd int) func(ctx znet.IrisCtx) {
// 	switch cmd {
// 	case apias.Cmd_AuthLogin:
// 		return apias.GetIrisHandle(&protbee.Login{}, func(h *zipbf.BeeHeader, req zproto.Message) (zproto.Message, error) {
// 			return a.BeeResLogin(h, req)
// 		})
// 	default:
// 		return nil
// 	}
// }

// func (a *ApiAuth) getIrisHandle(cmd int, req zproto.Message) func(ctx znet.IrisCtx) {
// 	switch cmd {
// 	case apias.Cmd_AuthLogin:
// 	default:
// 		return nil
// 	}
// }
