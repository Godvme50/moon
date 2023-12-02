package model

import (
	query "github.com/aide-cloud/gorm-normalize"
)

const TableNamePromNotify = "prom_alarm_notifies"

// PromAlarmNotify 告警通知对象
type PromAlarmNotify struct {
	query.BaseModel
	Name            string                   `gorm:"column:name;type:varchar(64);not null;uniqueIndex:idx__name,priority:1;comment:通知名称"`
	Status          int32                    `gorm:"column:status;type:tinyint;not null;default:1;comment:状态"`
	Remark          string                   `gorm:"column:remark;type:varchar(255);not null;comment:备注"`
	ChatGroups      []*PromAlarmChatGroup    `gorm:"many2many:prom_notify_chat_groups;comment:通知组"`
	BeNotifyMembers []*PromAlarmNotifyMember `gorm:"comment:被通知成员"`
}

func (*PromAlarmNotify) TableName() string {
	return TableNamePromNotify
}