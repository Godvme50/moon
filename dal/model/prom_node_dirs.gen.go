// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePromNodeDir = "prom_node_dirs"

// PromNodeDir mapped from table <prom_node_dirs>
type PromNodeDir struct {
	ID        int32              `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	NodeID    int32              `gorm:"column:node_id;type:int unsigned;not null;comment:节点 ID" json:"node_id"` // 节点 ID
	Path      *string            `gorm:"column:path;type:varchar(255);comment:目录地址" json:"path"`                 // 目录地址
	CreatedAt time.Time          `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt *time.Time         `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt int64              `gorm:"column:deleted_at;type:bigint unsigned;not null" json:"deleted_at"`
	Files     []*PromNodeDirFile `gorm:"foreignKey:DirID" json:"files"`
}

// TableName PromNodeDir's table name
func (*PromNodeDir) TableName() string {
	return TableNamePromNodeDir
}
