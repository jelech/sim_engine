package engine

import (
	"sync"
)

/*
StateMachine 状态机，用于管理实体的状态转换
根据不同的事件类型和当前状态，决定实体的下一个状态
并可能产生新的事件
*/

type EventHandler func(*Event) error

type StateMachine struct {
	transitions map[EventType]EventHandler // key: event_type, value: func() // 事件类型到状态转换函数的映射
	mu          sync.RWMutex
}

func NewStateMachine() *StateMachine {
	return &StateMachine{
		transitions: make(map[EventType]EventHandler),
	}
}

func (sm *StateMachine) Registry(eventType EventType, handler EventHandler) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.transitions[eventType] = handler
}
