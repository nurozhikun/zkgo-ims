package zidbauth

import (
	"testing"

	"gitee.com/sienectagv/gozk/zproto"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf/protbee"
)

func TestUserReq(t *testing.T) {
	// userReq := &protbee.UserReq{
	// 	Password: "123456",
	// 	User:     "admin",
	// }
	bs := []byte(`{
		"user": "admin",
		"password": "123456"
	}`)

	var msg zproto.Message
	msg = &protbee.UserReq{}
	// Unmarshal使用的是json格式
	if err := zproto.UnmarshalString(bs, msg); err != nil {
		t.Log(err)
	}

	userReq, ok := msg.(*protbee.UserReq)
	t.Log(ok)
	t.Log(userReq.User)
}
