package zidbauth

import (
	"gitee.com/sienectagv/gozk/zsql"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql/zidb"
)

type DB struct {
	*zidb.DB
}

func (db *DB) Migrate() error {
	return nil
}

func (db *DB) ZSqlDB() *zsql.DB {
	return db.DB
}
