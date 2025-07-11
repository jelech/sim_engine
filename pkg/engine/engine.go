package engine

import (
	"sync"
	"time"
)

/*
引擎是整个事件触发引擎的核心组件，它负责管理事件的注册和统一管理
*/
type Engine struct {
	WorldSystem *WorldSystem // 世界系统
	Running     bool         // 引擎运行状态
	StartTime   time.Time    // 引擎启动时间
	CurrentTime time.Time    // 当前仿真时间
	mu          sync.RWMutex // 保护并发访问
}

func NewEngine() *Engine {
	return &Engine{
		WorldSystem: NewWorldSystem(),
		Running:     false,
		StartTime:   time.Now(),
		CurrentTime: time.Now(),
	}
}

func (e *Engine) RegistryWorldSystem(ws *WorldSystem) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.WorldSystem = ws
}

func (e *Engine) Start() {
	if e.Running {
		return
	}

	e.Running = true
	e.StartTime = time.Now()
	e.CurrentTime = e.StartTime

	// 启动世界系统
	go e.WorldSystem.Start()
}

func (e *Engine) Stop() {
	if !e.Running {
		return
	}

	e.Running = false
	e.WorldSystem.Stop()
	e.CurrentTime = time.Now() // 更新当前时间为停止时刻
}
