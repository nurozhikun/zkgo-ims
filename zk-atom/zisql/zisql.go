package zisql

import (
	"gitee.com/sienectagv/gozk/zsql"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql/zidbauth"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql/zidbmap"
)

func WrapupDbMap(db *zsql.DB) *zidbmap.DB {
	return &zidbmap.DB{DB: db}
}

func WrapupDbAuth(db *zsql.DB) *zidbauth.DB {
	audb := &zidbauth.DB{}
	audb.DB = db
	zidbauth.SetJwtDB(audb)
	return audb
}

// type IProtBeeDB interface {
// 	ZSqlDB() *zsql.DB
// 	// ProtBeeResponse(h *protbee.Header, req zproto.Message) func(h *protbee.Header, req zproto.Message) (zproto.Message, error)
// }

// func BeeHandleByFuncName(db IProtBeeDB, method string) func(h *protbee.Header, req zproto.Message) (zproto.Message, error) {
// 	if nil == db {
// 		return nil
// 	}
// 	value := reflect.ValueOf(db)
// 	if !value.IsValid() {
// 		return nil
// 	}
// 	vf := value.MethodByName(method)
// 	if !vf.IsValid() {
// 		return nil
// 	}
// 	return func(h *protbee.Header, req zproto.Message) (zproto.Message, error) {
// 		in := make([]reflect.Value, 2)
// 		in[0] = reflect.ValueOf(h)
// 		in[1] = reflect.ValueOf(req)
// 		out := vf.Call(in)
// 		if len(out) < 2 {
// 			return nil, zutils.NewError(-1, "the response return is wrong")
// 		}
// 		msg, _ := out[0].Interface().(zproto.Message)
// 		err, _ := out[1].Interface().(error)
// 		return msg, err
// 	}
// }
