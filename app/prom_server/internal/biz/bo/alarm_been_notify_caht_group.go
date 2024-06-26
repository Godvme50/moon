package bo

import (
	"github.com/aide-family/moon/api"
	"github.com/aide-family/moon/app/prom_server/internal/biz/do"
	"github.com/aide-family/moon/app/prom_server/internal/biz/vobj"
)

type (
	PromAlarmBeenNotifyChatGroupBO struct {
		ID                uint32       `json:"id"`
		Status            vobj.Status  `json:"status"`
		Msg               string       `json:"msg"`
		PromAlarmNotifyID uint32       `json:"promAlarmNotifyID"`
		RealtimeAlarmID   uint32       `json:"realtimeAlarmID"`
		ChatGroup         *ChatGroupBO `json:"chatGroup"`
		ChatGroupId       uint32       `json:"chatGroupId"`

		CreatedAt int64 `json:"createdAt"`
		UpdatedAt int64 `json:"updatedAt"`
		DeletedAt int64 `json:"deletedAt"`
	}
)

// GetChatGroup .
func (l *PromAlarmBeenNotifyChatGroupBO) GetChatGroup() *ChatGroupBO {
	if l.ChatGroup == nil {
		return nil
	}
	return l.ChatGroup
}

// ToModel 转换为模型
func (l *PromAlarmBeenNotifyChatGroupBO) ToModel() *do.PromAlarmBeenNotifyChatGroup {
	return &do.PromAlarmBeenNotifyChatGroup{
		BaseModel:         do.BaseModel{ID: l.ID},
		RealtimeAlarmID:   l.RealtimeAlarmID,
		ChatGroup:         l.GetChatGroup().ToModel(),
		ChatGroupId:       l.ChatGroupId,
		Status:            l.Status,
		Msg:               l.Msg,
		PromAlarmNotifyID: l.PromAlarmNotifyID,
	}
}

// ToApi 转换为api对象
func (l *PromAlarmBeenNotifyChatGroupBO) ToApi() *api.ChatGroupSelectV1 {
	if l == nil {
		return nil
	}
	return &api.ChatGroupSelectV1{
		Value:  l.ID,
		App:    0,
		Label:  "",
		Status: l.Status.Value(),
	}
}

// PromAlarmBeenNotifyChatGroupModelToBO 转换为业务对象
func PromAlarmBeenNotifyChatGroupModelToBO(m *do.PromAlarmBeenNotifyChatGroup) *PromAlarmBeenNotifyChatGroupBO {
	return &PromAlarmBeenNotifyChatGroupBO{
		ID:                m.ID,
		Status:            m.Status,
		Msg:               m.Msg,
		PromAlarmNotifyID: m.PromAlarmNotifyID,
		RealtimeAlarmID:   m.RealtimeAlarmID,
		ChatGroup:         ChatGroupModelToBO(m.GetChatGroup()),
		ChatGroupId:       m.ChatGroupId,
		CreatedAt:         m.CreatedAt.Unix(),
		UpdatedAt:         m.UpdatedAt.Unix(),
		DeletedAt:         int64(m.DeletedAt),
	}
}
