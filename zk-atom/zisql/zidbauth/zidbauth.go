package zidbauth

import (
	"gitee.com/sienectagv/gozk/zlogger"
	"gitee.com/sienectagv/gozk/zsql"
	"gitee.com/sienectagv/gozk/zutils"
	"github.com/nurozhikun/zkgo-ims/sutils"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql/zidb"
)

type DB struct {
	*zidb.DB
}

func CreateDB(cfg *zsql.Cfg) *DB {
	db := &DB{DB: zsql.OpenDB(cfg)}
	if db.DB == nil {
		zlogger.Error("create db failed")
		return nil
	}
	if err := db.Migrate(); err != nil {
		zlogger.Error(err)
	}
	return db
}

func (db *DB) DeleteDB() {
	if db.DB == nil {
		return
	}
	db.DB.Close()
}

func (db *DB) Migrate() error {
	err := db.AutoMigrate(&User{}, &Role{}, &Permission{})
	if err != nil {
		return err
	}
	_, err = db.AddForeignKey("users", "editor_id", "users", "id")
	if err != nil {
		return err
	}
	_, err = db.AddForeignKey("permissions", "parent_id", "permissions", "id")
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) ZSqlDB() *zsql.DB {
	return db.DB
}

func (db *DB) CheckInUser(user, password string) (*User, error) {
	row := &User{}
	result := db.Where("name = ?", user).Find(row)
	if result.RowsAffected == 0 {
		return nil, zutils.ErrUserOrPassMiss
	}

	if err := db.Model(row).Association("Roles").Find(&(row.Roles)); err != nil {
		zlogger.Error(err)
		return nil, err
	}

	err := sutils.SaltPassCheck(user, password, row.Salt, row.Password)
	if nil != err {
		return nil, err
	}
	return row, nil
}
