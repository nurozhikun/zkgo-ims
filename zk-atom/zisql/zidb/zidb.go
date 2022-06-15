package zidb

import (
	"gitee.com/sienectagv/gozk/zsql"
	"gorm.io/gorm"
)

type Model = gorm.Model

type Entity struct {
	Code      string `gorm:"type:varchar(64); unique; not null;"`
	Name      string `gorm:"type:varchar(128);not null;default:''"`
	LifeCycle int32  `gorm:"not null;default:0;"`
}

type DB = zsql.DB

type Base struct {
	*DB
}

func (db *Base) ZSqlDB() *zsql.DB {
	return db.DB
}

// func (db *DB) ZSqlDB() *zsql.DB {
// 	return (*zsql.DB)(db)
// }

// type DB struct {
// 	*zsql.DB
// }

// type ddb struct {
// 	DB
// }

// func NewDb(d *zsql.DB) *DB {
// 	return &DB{DB: d}
// }
