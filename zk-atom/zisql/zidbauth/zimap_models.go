package zidbauth

import "github.com/nurozhikun/zkgo-ims/zk-atom/zisql/zidb"

type User struct {
	// ID    uint `gorm:"primarykey;autoIncrement;"`
	Model      zidb.Model `gorm:"embedded"`
	State      int32      `gorm:"type:tinyint(1);not null; default:0; comment:'0=permitted,1=useful';"`
	Name       string     `gorm:"type:varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci; unique;not null;"`
	NickName   string     `gorm:"type:varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci;"`
	HeadImgUrl string     `gorm:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci;"`
	Mobile     string     `gorm:"type:varchar(12) CHARACTER SET utf8 COLLATE utf8_general_ci;"`
	Salt       string     `gorm:"type:varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci;"`
	Password   []byte     //`gorm:"type:varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci;"`
	EditorID   *uint      // self-referential has one; maybe NULL,
	Editor     *User      //`gorm:"foreignKey:EditorID; constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`
	Roles      []*Role    `gorm:"many2many:user_roles"`
}

type Account struct {
	Model    zidb.Model `gorm:"embedded"`
	OpenCode string     `gorm:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci;not null;"`
	Category int32      `gorm:"not null; default 0; comment:'0表示account就是用户名';"`
	// Deleted  int32  `gorm:"type:tinyint(1);not null;default 0;"`
	UserID uint
	User   User
}

type Role struct {
	Model zidb.Model `gorm:"embedded;"`
	Code  string     `gorm:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci;not null;"`
	Name  string     `gorm:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci;"`
	Intro string     `gorm:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci;"`
	// Deleted     int32  `gorm:"type:tinyint(1);not null;default 0;"`
	Level       int32 `gorm:"not null; default 0;"` // 简单模式可以只用level来控制权限
	ParentID    *uint
	Parent      *Role        //`gorm:"foreignkey:ParentID; constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
}

type Permission struct {
	Model    zidb.Model `gorm:"embedded;"`
	Url      string     `gorm:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci; unique; not null;"`
	Name     string     `gorm:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci; unique;"`
	Intro    string     `gorm:"type:varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci;"`
	Category int32      `gorm:"type:tinyint;not null; default:0;"`
	// Deleted  int32  `gorm:"type:tinyint(1);not null;default 0;"`
	ParentID *uint
	Parent   *Permission `gorm:"foreignkey:ParentID;constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`
}
