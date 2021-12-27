package zisql

import (
	"gitee.com/sienectagv/gozk/zsql"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql/zidbmap"
)

func WrapupDbMap(db *zsql.DB) *zidbmap.DB {
	return &zidbmap.DB{DB: db}
}
