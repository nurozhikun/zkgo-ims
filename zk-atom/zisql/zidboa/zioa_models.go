package zidboa

import "github.com/nurozhikun/zkgo-ims/zk-atom/zisql/zidb"

type Group struct {
	Model  zidb.Model  `gorm:"embedded;"`
	Entity zidb.Entity `gorm:"embedded;"`
	// Code  string     `gorm:"type:varchar(32);unique;not null;"`
	// Alias string     `gorm:"type:varchar(64);"`
}

//unique(Code, GroupID)
type Factory struct {
	Model     zidb.Model `gorm:"embedded;"`
	Code      string     `gorm:"type:varchar(64); not null;"`
	Name      string     `gorm:"type:varchar(128);not null;default:''"`
	LifeCycle int32      `gorm:"not null;default:0;"`
	GroupID   uint       //foreign key to Group
	Group     Group
}
