package zidbdevice

import "github.com/nurozhikun/zkgo-ims/zk-atom/zisql/zidb"

type Device struct {
	Model   zidb.Model `gorm:"embedded;"`
	Code    string     `gorm:"type:varchar(32);not null;uniqueindex:code;"`
	GroupID uint       `gorm:"not null;uniqueindex:code"`
}
