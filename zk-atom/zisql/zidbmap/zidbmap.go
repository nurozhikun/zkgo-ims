package zidbmap

import (
	"gitee.com/sienectagv/gozk/zsql"
)

type DB struct {
	*zsql.DB
}

func (db *DB) Migrate() error {
	err := db.AutoMigrate(&Map{}, &Point{}, &Line{})
	return err
}

func (db *DB) ZSqlDB() *zsql.DB {
	return db.DB
}
