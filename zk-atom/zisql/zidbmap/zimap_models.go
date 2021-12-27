package zidbmap

import "github.com/nurozhikun/zkgo-ims/zk-atom/zisql/zidb"

type Map struct {
	Model          zidb.Model `gorm:"embedded;"`
	Code           string     `gorm:"type:varchar(32);unique;not null;"`
	Alias          string     `gorm:"type:varchar(64);"`
	Version        string     `gorm:"type:varchar(32);not null;"`
	LifeCycle      int32      `gorm:"not null;default 0;"` //disable 0; working 1; editing 2; 3 uploading; 4 downloading;
	EditingVersion string     `gorm:"type:varchar(32);"`   //正在编辑的版本，如果是null，表示没在编辑；一个地图同时只能从一台设备编辑
	Left           float32
	Right          float32
	Top            float32
	Bottom         float32
}
type Point struct {
	Model     zidb.Model `gorm:"embedded;"`
	Type      int32      `gorm:"not null;default:0;"`
	Code      string     `gorm:"type:varchar(32);unique;not null;"`
	MapID     uint       // foreign key to Map
	Map       Map
	LifeCycle int32 `gorm:"not null;default:0;"`
	PosX      float32
	PosY      float32
	Lines     []Line `gorm:"many2many:point_lines;"` // 会创建r_point_lines关联表
}

type Line struct {
	Model        zidb.Model `gorm:"embedded;"`
	ToPointID    uint       // use foreignkey
	ToPoint      Point      `gorm:"foreignkey:ToPointID;"` // the column 'ToPointID' is a foreign key to the Table 'Point'
	Distance     float32    `gorm:"not null;default:0.0;"`
	GeneralActor float32    `gorm:"not null;default:1.0;"`
	// Weights   []*Weight `gorm:"many2many:line_weights;"` //会创建r_line_weights关联表
}

//激光点存储，需要分区块
type LaserBlock struct {
	Model  zidb.Model `gorm:"embedded;"`
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
	MapID  uint64
	Map    Map
}

type LaserSpot struct {
	ID           uint `gorm:"primarykey;autoIncrement;"`
	X            int32
	Y            int32
	Z            int32
	Confidence   int32
	LaserBlockID uint
	LaserBlock   LaserBlock
}

//记录原始的激光数据
type LaserRaw struct {
	ID      uint `gorm:"primarykey;autoIncrement;"`
	Created int64
	Spots   []byte
}
