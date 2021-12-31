package zidbauth

import (
	"gitee.com/sienectagv/gozk/zproto"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf/protbee"
)

func (db *DB) ProtBeeResponse(cmd int, h *protbee.Header, res zproto.Message) {

}

func (db *DB) BeeMapLogin(header *zipbf.BeeHeader, reqBody zproto.Message) (zproto.Message, error) {
	body := &protbee.LoginAck{}
	//TODO
	return body, nil
}
