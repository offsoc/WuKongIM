package types

import (
	"context"

	"github.com/WuKongIM/WuKongIM/internal/types/pluginproto"
	wkproto "github.com/WuKongIM/WuKongIMGoProto"
)

type PluginStatus uint32

const (
	PluginStatusInit     PluginStatus = iota
	PluginStatusNormal                = 1 // 插件正常
	PluginStatusError                 = 2 // 插件异常
	PluginStatusOffline               = 3 // 插件离线
	PluginStatusDisabled              = 4 // 插件禁用
)

func (s PluginStatus) String() string {
	switch s {
	case PluginStatusInit:
		return "init"
	case PluginStatusNormal:
		return "normal"
	case PluginStatusError:
		return "error"
	case PluginStatusOffline:
		return "offline"
	case PluginStatusDisabled:
		return "disabled"
	}
	return "unknown"
}

type Plugin interface {
	// GetNo 获取插件编号
	GetNo() string
	// Send 调用插件的Send方法
	Send(ctx context.Context, sendPacket *pluginproto.SendPacket) (*pluginproto.SendPacket, error)
	// PersistAfter 调用插件的PersistAfter方法
	PersistAfter(ctx context.Context, messages *pluginproto.MessageBatch) error
	// Reply 调用插件的Reply方法
	Receive(ctx context.Context, recv *pluginproto.RecvPacket) error
	// Route 调用插件的Route方法
	Route(ctx context.Context, request *pluginproto.HttpRequest) (*pluginproto.HttpResponse, error)
	// Status 获取插件状态
	Status() PluginStatus
}

type PluginMethod string

const (
	PluginSend         PluginMethod = "Send"
	PluginPersistAfter PluginMethod = "PersistAfter"
	PluginReceive      PluginMethod = "Receive"
	PluginRoute        PluginMethod = "Route"
	PluginConfigUpdate PluginMethod = "ConfigUpdate"
)

func (p PluginMethod) String() string {
	return string(p)
}

type PluginMethodType uint32

const (
	PluginMethodTypeSend         PluginMethodType = 1
	PluginMethodTypePersistAfter PluginMethodType = 2
	PluginMethodTypeReceive      PluginMethodType = 3
	PluginMethodTypeRoute        PluginMethodType = 4
	PluginMethodTypeConfigUpdate PluginMethodType = 5
)

func (p PluginMethod) Type() PluginMethodType {
	switch p {
	case PluginSend:
		return PluginMethodTypeSend
	case PluginPersistAfter:
		return PluginMethodTypePersistAfter
	case PluginReceive:
		return PluginMethodTypeReceive
	case PluginRoute:
		return PluginMethodTypeRoute
	case PluginConfigUpdate:
		return PluginMethodTypeConfigUpdate
	}
	return 0
}

type PluginResponse struct {
	MessageId uint64
	Frame     wkproto.Frame
}
