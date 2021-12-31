package test

import (
	"reflect"
	"testing"

	"gitee.com/sienectagv/gozk/zlogger"
	"gitee.com/sienectagv/gozk/zproto"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi/apias"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf/protbee"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql"
)

func TestInstallHandler(t *testing.T) {
	// apis := &ziapi.IrisBeeApis{}
	// apis.InstallApi(ziapi.NewApiMap(nil))
	api := ziapi.NewApiMap(nil)
	name, ok := zisql.BeeFuncNames[apias.Cmd_MapThumbnails]
	fn := ziapi.BeeHandleByFuncName(api, name)
	zlogger.Info(ok, reflect.TypeOf(fn), reflect.ValueOf(fn))
	m, err := fn(&protbee.Header{}, &zproto.EmptyMessage{})
	zlogger.Info(m, err)
}
