package engine

import (
	"sync"
	"time"
)

/*
世界系统是引擎的核心组件之一，它负责管理世界的状态和事件
其包含多个区域, 每个区域可以有自己的状态和事件;
其同步多个区域之间的状态和事件;
其支持对外部触发的事件进行转发;
其支持自动触发时间同步时间, 让区域时间同步一致;
其拥有独立的优先队列;
*/
type WorldSystem struct {
	ID          string                   // 世界系统唯一标识
	Regions     map[string]*RegionSystem // 管理的区域系统
	EventQueue  *PriorityEventQueue      // 世界级别的事件优先队列
	CurrentTime time.Time                // 当前世界时间
	TimeStep    time.Duration            // 时间步长
	Running     bool                     // 运行状态
	mu          sync.RWMutex             // 保护并发访问
	stopChan    chan struct{}            // 停止信号通道
}

func NewWorldSystem() *WorldSystem {
	return &WorldSystem{
		ID:          "world",
		Regions:     make(map[string]*RegionSystem),
		EventQueue:  NewPriorityEventQueue(),
		CurrentTime: time.Now(),
		TimeStep:    time.Millisecond * 100, // 默认100ms步长
		Running:     false,
		stopChan:    make(chan struct{}),
	}
}

func (ws *WorldSystem) RegistryRegionSystem(region *RegionSystem) {
	ws.Regions[region.ID] = region
	region.WorldSystem = ws // 设置区域系统的父级世界系统引用
}

func (ws *WorldSystem) RegistryRegionSystems(regions []*RegionSystem) {
	for _, region := range regions {
		ws.Regions[region.ID] = region
		region.WorldSystem = ws // 设置区域系统的父级世界系统引用
	}
}

func (ws *WorldSystem) PublishEvent(event *Event) {
	ws.EventQueue.Push(event)
}

func (ws *WorldSystem) Start() {
	if ws.Running {
		return
	}
	ws.Running = true

	go ws.processEvents()
}

func (ws *WorldSystem) Stop() {
	if !ws.Running {
		return
	}
	ws.Running = false

	close(ws.stopChan) // 发送停止信号
	for _, region := range ws.Regions {
		region.Stop() // 停止所有区域系统
	}
}

func (ws *WorldSystem) BroadcastEvent(event *Event) {
	for _, region := range ws.Regions {
		region.ReceiveEvent(event) // 将事件发送到每个区域系统
	}
}

func (ws *WorldSystem) processEvents() {
	for ws.Running {
		select {
		case <-ws.stopChan:
			return
		default:
			if eventItf := ws.EventQueue.Pop(); eventItf != nil {
				event, ok := eventItf.(*Event)
				if !ok {
					// 处理类型断言失败的情况
					continue
				}

				if region, exists := ws.Regions[event.TargetRegion]; exists {
					region.ReceiveEvent(event)
				} else {
					// 处理目标区域不存在的情况
					continue
				}
			}
		}
	}
}
