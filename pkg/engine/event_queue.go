package engine

import (
	"container/heap"
	"sync"
	"time"
)

/*
PriorityEventQueue 是基于优先级的事件队列
使用 Go 标准库的 heap 实现优先队列
事件按照优先级和调度时间进行排序
*/
type PriorityEventQueue struct {
	events []*Event
	mu     sync.RWMutex
}

var _ heap.Interface = (*PriorityEventQueue)(nil)

func NewPriorityEventQueue() *PriorityEventQueue {
	pq := &PriorityEventQueue{
		events: make([]*Event, 0),
	}
	heap.Init(pq)
	return pq
}

// Len 返回队列长度
func (pq *PriorityEventQueue) Len() int {
	return len(pq.events)
}

// Less 比较两个事件的优先级
// 优先级高的排在前面，优先级相同时按调度时间排序
func (pq *PriorityEventQueue) Less(i, j int) bool {
	// 先比较优先级（优先级高的排前面）
	if pq.events[i].Priority != pq.events[j].Priority {
		return pq.events[i].Priority > pq.events[j].Priority
	}
	// 优先级相同时，调度时间早的排前面
	return pq.events[i].ScheduleTime.Before(pq.events[j].ScheduleTime)
}

// Swap 交换两个元素
func (pq *PriorityEventQueue) Swap(i, j int) {
	pq.events[i], pq.events[j] = pq.events[j], pq.events[i]
}

// Push 添加元素到队列
func (pq *PriorityEventQueue) Push(x interface{}) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	pq.events = append(pq.events, x.(*Event))
}

// Pop 从队列中移除并返回优先级最高的元素
func (pq *PriorityEventQueue) Pop() interface{} {
	old := pq.events
	n := len(old)
	event := old[n-1]
	pq.events = old[0 : n-1]
	return event
}

// Enqueue 线程安全地添加事件到队列
func (pq *PriorityEventQueue) Enqueue(event *Event) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	heap.Push(pq, event)
}

// Dequeue 线程安全地从队列中取出优先级最高的事件
func (pq *PriorityEventQueue) Dequeue() *Event {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	if pq.Len() == 0 {
		return nil
	}

	return heap.Pop(pq).(*Event)
}

// Peek 查看队列头部的事件但不移除
func (pq *PriorityEventQueue) Peek() *Event {
	pq.mu.RLock()
	defer pq.mu.RUnlock()

	if pq.Len() == 0 {
		return nil
	}

	return pq.events[0]
}

// Size 返回队列大小
func (pq *PriorityEventQueue) Size() int {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	return pq.Len()
}

// IsEmpty 检查队列是否为空
func (pq *PriorityEventQueue) IsEmpty() bool {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	return pq.Len() == 0
}

// Clear 清空队列
func (pq *PriorityEventQueue) Clear() {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	pq.events = pq.events[:0]
	heap.Init(pq)
}

// GetEventsReadyAt 获取指定时间点之前应该执行的所有事件
func (pq *PriorityEventQueue) GetEventsReadyAt(currentTime time.Time) []*Event {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	var readyEvents []*Event

	// 从队列头部开始检查事件
	for pq.Len() > 0 {
		event := pq.events[0]
		if event.ScheduleTime.After(currentTime) {
			break // 如果头部事件的调度时间还没到，后面的也不会到
		}

		readyEvents = append(readyEvents, heap.Pop(pq).(*Event))
	}

	return readyEvents
}
