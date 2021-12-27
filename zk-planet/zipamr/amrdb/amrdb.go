package amrdb

import (
	"gitee.com/sienectagv/gozk/zsql"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql"
)

type DB struct {
	*zsql.DB
}

func CreateDB(cfg *zsql.Cfg) *DB {
	db := &DB{DB: zsql.OpenDB(cfg)}
	if nil != db.DB {
		db.init()
	}
	return db
}

func (db *DB) DeleteDB() {
	if nil == db.DB {
		return
	}
	db.DB.Close()
}

func (db *DB) init() {
	zisql.WrapupDbMap(db.DB).Migrate()
}
