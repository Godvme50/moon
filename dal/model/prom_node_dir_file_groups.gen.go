// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePromNodeDirFileGroup = "prom_node_dir_file_groups"

// PromNodeDirFileGroup mapped from table <prom_node_dir_file_groups>
type PromNodeDirFileGroup struct {
	ID         int32                           `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	Name       string                          `gorm:"column:name;type:varchar(64);not null;default:alert_;comment:策略组名称" json:"name"` // 策略组名称
	Remark     string                          `gorm:"column:remark;type:varchar(255);not null;comment:说明" json:"remark"`              // 说明
	CreatedAt  time.Time                       `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt  *time.Time                      `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt  int64                           `gorm:"column:deleted_at;type:bigint unsigned;not null" json:"deleted_at"`
	FileID     int32                           `gorm:"column:file_id;type:int unsigned;not null;comment:files id" json:"file_id"` // files id
	Strategies []*PromNodeDirFileGroupStrategy `gorm:"foreignKey:GroupID" json:"strategies"`
}

// TableName PromNodeDirFileGroup's table name
func (*PromNodeDirFileGroup) TableName() string {
	return TableNamePromNodeDirFileGroup
}
