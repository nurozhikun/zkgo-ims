package ziapp

import (
	"gitee.com/sienectagv/gozk/zsql"
	"github.com/kataras/iris/v12"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql"
)

type ZiAppBee struct {
	*iris.Application
	SqlDbs
	ziapi.HttpHandles
}

func (ap *ZiAppBee) Init(App *iris.Application) {
	ap.SqlDbs = make(SqlDbs)
	ap.Application = App
}

//for default api of auth
var DefAuthPath = "/auth"

func (ap *ZiAppBee) InstallBeeProtAuthDefault(cfg *zsql.Cfg) error {
	ap.AppendDb(DbKeyAuth, cfg)
	db, err := ap.DbOf(DbKeyAuth)
	if err != nil {
		return err
	}
	ap.AppendHandle(zisql.WrapupDbAuth(db))
	party := ap.Party(DefAuthPath)
	ap.InstallBeeProtobufOfCmds(party, ziapi.AuthBeeCmds...)
	return nil
}
