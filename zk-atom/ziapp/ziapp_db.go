package ziapp

import (
	"fmt"

	"gitee.com/sienectagv/gozk/zsql"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql"
)

const (
	DbKeyAuth = "auth"
)

type SqlDbs map[string]*zsql.DB

func (dbs *SqlDbs) AppendDb(key string, cfg *zsql.Cfg) {
	db := zsql.OpenDB(cfg)
	if DbKeyAuth == key {
		zisql.WrapupDbAuth(db).Migrate()
	}
	(*dbs)[key] = db
}

func (dbs *SqlDbs) DbOf(key string) (*zsql.DB, error) {
	db, ok := (*dbs)[key]
	if !ok {
		return nil, fmt.Errorf("the db key %s is not exist", key)
	}
	return db, nil
}
