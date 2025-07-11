package engine

import (
	"sync"
	"time"
)

/*
区域系统是世界系统的子系统之一，它负责管理区域的状态和事件
其包含多个实体, 每个实体可以有自己的状态和事件;
多个实体之间几乎没有交互, 但可以通过事件进行通信;
多个区域之间可以通过世界系统事件进行通信;
其拥有独立的优先队列;
其拥有一套状态机, 用于管理实体的状态管理, 产生新的事件;
*/
// 作用: 管理区域内的实体和事件，处理区域内的状态机逻辑
type RegionSystem struct {
	ID           string              // 区域唯一标识
	Name         string              // 区域名称
	Entities     map[string]*Entity  // 管理的实体
	EventQueue   *PriorityEventQueue // 区域级别的事件优先队列
	StateMachine *StateMachine       // 状态机，用于处理实体状态转换
	WorldSystem  *WorldSystem        // 引用父级世界系统
	CurrentTime  time.Time           // 当前区域时间
	Running      bool                // 运行状态
	mu           sync.RWMutex        // 保护并发访问
	stopChan     chan struct{}       // 停止信号通道
	timeStep     time.Duration       // 时间步长
}

func NewRegionSystem(id, name string, worldSystem *WorldSystem) *RegionSystem {
	return &RegionSystem{
		ID:           id,
		Name:         name,
		Entities:     make(map[string]*Entity),
		EventQueue:   NewPriorityEventQueue(),
		StateMachine: NewStateMachine(),
		WorldSystem:  worldSystem,
		CurrentTime:  time.Now(),
		Running:      false,
		stopChan:     make(chan struct{}),
	}
}

func (rs *RegionSystem) RegistryEntities(entities []*Entity) {
	for _, entity := range entities {
		rs.Entities[entity.ID] = entity
	}
}

func (rs *RegionSystem) ApplyEvent(event *Event) {
	if handler, exists := rs.StateMachine.transitions[event.Type]; exists {
		handler(event)
	}
}

func (rs *RegionSystem) ReceiveEvent(event *Event) {
	rs.EventQueue.Push(event)
}

func (rs *RegionSystem) GenerateEvent(event *Event) {
	if event.TargetRegion == "" || event.TargetRegion == rs.ID {
		rs.EventQueue.Push(event) // 将事件添加到区域事件队列
		return
	}

	// 如果事件指定了目标区域，则发送到世界系统进行转发
	if rs.WorldSystem != nil {
		rs.WorldSystem.PublishEvent(event) // 发布事件到世界系统
	}
}

func (rs *RegionSystem) BroadcastEvent(event *Event) {
	for _, entity := range rs.Entities {
		entity.mu.Lock()
		entity.Properties["last_event"] = event // 更新实体的最后事件属性
		entity.mu.Unlock()
	}
}

func (rs *RegionSystem) Start() {
	if rs.Running {
		return
	}
	rs.Running = true

	go rs.processLocalEvents()
}

func (rs *RegionSystem) Stop() {
	if !rs.Running {
		return
	}
	rs.Running = false
	close(rs.stopChan)
}

func (rs *RegionSystem) processLocalEvents() {
	for {
		time.Sleep(rs.timeStep) // 模拟时间步进
		rs.CurrentTime = rs.CurrentTime.Add(rs.timeStep)

		select {
		case <-rs.stopChan:
			return
		default:
			peekEvent := rs.EventQueue.Peek()
			if peekEvent.ScheduleTime.After(rs.CurrentTime) {
				continue
			}

			if event := rs.EventQueue.Dequeue(); event != nil {
				rs.ApplyEvent(event)
				rs.CurrentTime = peekEvent.ScheduleTime
			}
		}
	}
}
