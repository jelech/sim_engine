package engine

import (
	"sync"
	"time"
)

// EntityType 定义实体类型
type EntityType string

/*
Entity 是引擎中的实体对象，代表仿真世界中的一个具体对象
实体可以是仓库、SKU、货架、工人等，它们可以拥有自己的状态和属性
实体可以通过事件进行通信和交互
实体可以被区域系统管理，并参与到区域的状态和事件中
*/
type EntityItf interface {
	ID() string       // 实体唯一标识
	Type() EntityType // 实体类型
	Name() string     // 实体名称
	RegionID() string // 所属区域ID
}

type Entity struct {
	ID         string                 // 实体唯一标识
	Type       EntityType             // 实体类型
	Name       string                 // 实体名称
	RegionID   string                 // 所属区域ID
	Properties map[string]interface{} // 实体属性
	CreateTime time.Time              // 创建时间
	UpdateTime time.Time              // 最后更新时间
	mu         sync.RWMutex           // 保护并发访问
}

func NewEntity(id string, entityType EntityType, name, regionID string) *Entity {
	now := time.Now()
	return &Entity{
		ID:         id,
		Type:       entityType,
		Name:       name,
		RegionID:   regionID,
		Properties: make(map[string]interface{}),
		CreateTime: now,
		UpdateTime: now,
	}
}
