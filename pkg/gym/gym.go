// Package gym provides OpenAI Gym-compatible interface for reinforcement learning
package gym

type EngineGym interface {
	Reset() (observation interface{}, err error) // 重置环境，返回初始观测
	Step(action interface{}) (observation interface{}, reward float64, done bool, err error)
}

type EngineTask interface {
	Start() (metrics interface{}, err error) // 启动任务
}
