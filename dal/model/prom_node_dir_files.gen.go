// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePromNodeDirFile = "prom_node_dir_files"

// PromNodeDirFile mapped from table <prom_node_dir_files>
type PromNodeDirFile struct {
	ID        int32                   `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	Filename  string                  `gorm:"column:filename;type:varchar(64);not null;comment:yaml file" json:"filename"` // yaml file
	CreatedAt time.Time               `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt *time.Time              `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt int64                   `gorm:"column:deleted_at;type:bigint unsigned;not null" json:"deleted_at"`
	DirID     int32                   `gorm:"column:dir_id;type:int unsigned;not null;comment:目录 ID" json:"dir_id"` // 目录 ID
	Groups    []*PromNodeDirFileGroup `gorm:"foreignKey:FileID" json:"groups"`
}

// TableName PromNodeDirFile's table name
func (*PromNodeDirFile) TableName() string {
	return TableNamePromNodeDirFile
}
