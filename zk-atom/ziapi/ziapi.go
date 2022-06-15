package ziapi

import (
	"gitee.com/sienectagv/gozk/zproto"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf/protbee"
)

type beeResponseMethod = func(h *protbee.Header, req zproto.Message) (zproto.Message, error)

// type IProtoBeeHandle interface {
// 	// // the method of cmd must be func(h *protbee.Header, req zproto.Message) (zproto.Message, error)
// 	// MethodNameOfCmd(cmd int) (string, bool)
// 	// PathOfCmd(cmd int) (string, bool)
// 	// ReqBodyOfCmd(cmd int) zproto.Message
// }

type httpHanle interface{}

type HttpHandles struct {
	handles []httpHanle
}

func (iba *HttpHandles) AppendHandle(h httpHanle) {
	iba.handles = append(iba.handles, h)
}
