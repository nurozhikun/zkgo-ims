package zidbauth

import (
	"gitee.com/sienectagv/gozk/zlogger"
	"gitee.com/sienectagv/gozk/zsql"
	"gitee.com/sienectagv/gozk/zutils"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql/zidb"
)

type DB struct {
	// *zidb.DB
	zidb.Base
}

func (db *DB) ZSqlDB() *zsql.DB {
	return db.DB
}

// 保存一个authDB 如果此处不预留authDB 则需要从接口入手

func CreateDB(cfg *zsql.Cfg) *DB {
	// db := &DB{DB: zsql.OpenDB(cfg)}
	// db := &DB{DB: zsql.OpenDB(cfg)}
	db := &DB{}
	db.DB = zsql.OpenDB(cfg)
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

func (db *DB) CheckInUser(username, password string) (*User, error) {
	user := &User{}
	result := db.Where("name = ?", username).Find(user)
	if result.RowsAffected == 0 {
		return nil, zutils.ErrUserOrPassMiss
	}

	// 关联模式
	if err := db.Model(user).Association("Roles").Find(&(user.Roles)); err != nil {
		zlogger.Error(err)
		return nil, err
	}

	// 填充permissions
	for _, v := range user.Roles {
		if err := db.Model(v).Association("Permissions").Find(&(v.Permissions)); err != nil {
			zlogger.Error(err)
			return nil, err
		}
	}
	err := zidb.SaltPassCheck(username, password, user.Salt, user.Password)
	if nil != err {
		return nil, err
	}
	return user, nil
}
