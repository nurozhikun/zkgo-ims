package zidbtask

import "time"

//预约任务，设备切换，预约搬运
type TaskAppointment struct {
	ID          uint      `gorm:"primarykey;autoIncrement;comment:'id of the task;'"`
	AppointTime time.Time `gorm:"type:timestamp;not null;default:0;comment:'预约执行时间';"` //UTC.UnixMilli()
	Body        string    `gorm:"type:text;comment:'任务数据,json串';"`
	DeviceID    uint
}
