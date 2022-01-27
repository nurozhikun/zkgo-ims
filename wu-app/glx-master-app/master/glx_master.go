package master

import (
	"gitee.com/sienectagv/gozk/zsql"
	"github.com/kataras/iris/v12"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi"
)

type GlxMaster struct {
	App *iris.Application
	dbs map[string]*zsql.DB
	ziapi.IrisBeeApis
}

func NewGlxMaster() *GlxMaster {
	return &GlxMaster{
		App: iris.New(),
		dbs: make(map[string]*zsql.DB),
	}
}

func (gm *GlxMaster) AppendDb(name string, cfg *zsql.Cfg) {
	gm.dbs[name] = zsql.OpenDB(cfg)
}

func (gm *GlxMaster) DB(name string) *zsql.DB {
	db, _ := gm.dbs[name]
	return db
}
