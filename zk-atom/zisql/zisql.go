package zisql

import (
	"reflect"

	"gitee.com/sienectagv/gozk/zproto"
	"gitee.com/sienectagv/gozk/zsql"
	"gitee.com/sienectagv/gozk/zutils"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf/protbee"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql/zidb"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql/zidbauth"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql/zidbmap"
)

var (
	BeeFuncNames = zidb.BeeFuncNames
)

func WrapupDbMap(db *zsql.DB) *zidbmap.DB {
	return &zidbmap.DB{DB: db}
}

func WrapupDbAuth(db *zsql.DB) *zidbauth.DB {
	return &zidbauth.DB{DB: db}
}

type IProtBeeDB interface {
	ZSqlDB() *zsql.DB
	// ProtBeeResponse(h *protbee.Header, req zproto.Message) func(h *protbee.Header, req zproto.Message) (zproto.Message, error)
}

func BeeHandleByFuncName(db IProtBeeDB, method string) func(h *protbee.Header, req zproto.Message) (zproto.Message, error) {
	if nil == db {
		return nil
	}
	value := reflect.ValueOf(db)
	if !value.IsValid() {
		return nil
	}
	vf := value.MethodByName(method)
	if !vf.IsValid() {
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
