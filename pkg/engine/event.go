package engine

import (
	"time"
)

// EventType 定义事件类型
type EventType string

// EventPriority 定义事件优先级
type EventPriority int

/*
Event 是引擎中的事件对象，代表仿真世界中的一个具体事件
事件可以是实体的动作、状态变化、业务操作等
*/
type Event struct {
	ID           string                 // 事件唯一标识
	Type         EventType              // 事件类型
	Priority     EventPriority          // 事件优先级
	SourceRegion string                 // 事件源区域ID
	TargetRegion string                 // 事件目标区域ID
	SourceID     string                 // 事件源ID（可以是实体ID、区域ID等）
	TargetID     string                 // 事件目标ID
	Timestamp    time.Time              // 事件发生时间
	ScheduleTime time.Time              // 事件调度时间
	Data         map[string]interface{} // 事件数据
	Processed    bool                   // 是否已处理
}

func NewEvent(eventType EventType, sourceID, targetID string, data map[string]interface{}) *Event {
	now := time.Now()
	return &Event{
		ID:           generateEventID(),
		Type:         eventType,
		Priority:     1e10,
		SourceID:     sourceID,
		TargetID:     targetID,
		Timestamp:    now,
		ScheduleTime: now,
		Data:         data,
		Processed:    false,
	}
}

// generateEventID 生成唯一的事件ID
func generateEventID() string {
	return time.Now().Format("20060102150405") + "-" + randomString(6)
}

// randomString 生成指定长度的随机字符串
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(result)
}
