package engine_test

import (
	"testing"
	"time"

	"github.com/jelech/sim_engine/pkg/engine"
)

const (
	EventTypeEvent1 engine.EventType = "event1"
	EventTypeEvent2 engine.EventType = "event2"
)

type SelfRegionSystem struct {
	engine.RegionSystem
}

func NewSelfRegionSystem(id, name string, world *engine.WorldSystem) *SelfRegionSystem {
	return &SelfRegionSystem{
		RegionSystem: *engine.NewRegionSystem(id, name, world),
	}
}

func (s *SelfRegionSystem) EventTransfer_One(event *engine.Event) error {
	// 模拟事件处理逻辑
	s.CurrentTime = s.CurrentTime.Add(1 * time.Second)
	s.EventQueue.Push(&engine.Event{
		ID:           "test-event",
		Type:         EventTypeEvent1,
		SourceRegion: s.ID,
		TargetRegion: s.ID,
		Timestamp:    s.CurrentTime,
	})
	return nil
}
func (s *SelfRegionSystem) EventTransfer_Two(event *engine.Event) error {
	s.Entities[event.SourceID].Properties["msg"] = "running event transfer two"
	s.EventQueue.Push(&engine.Event{
		ID:           "test-event",
		Type:         EventTypeEvent1,
		SourceRegion: s.ID,
		TargetRegion: s.ID,
		Timestamp:    s.CurrentTime,
	})
	return nil
}

func TestCreateNewRegion(t *testing.T) {
	world := engine.NewWorldSystem()
	region := NewSelfRegionSystem("region-001", "区域A", world)
	entity1 := engine.NewEntity("entity-001", "type1", "实体1", "region-001")
	entity2 := engine.NewEntity("entity-002", "type2", "实体2", "region-001")

	region.RegistryEntities([]*engine.Entity{entity1, entity2})
	region.StateMachine.Registry(EventTypeEvent1, region.EventTransfer_One)
	region.StateMachine.Registry(EventTypeEvent2, region.EventTransfer_Two)

	// 模拟事件触发
	region.ApplyEvent(&engine.Event{
		ID:       "event-001",
		Type:     EventTypeEvent1,
		SourceID: entity1.ID,
		TargetID: entity2.ID,
	})

	// 输出区域系统信息
	println("区域系统ID:", region.ID)
	println("区域系统名称:", region.Name)
	println("实体数量:", len(region.Entities))
	for _, entity := range region.Entities {
		println(entity.Properties)
	}
}

func TestNewEngine(t *testing.T) {
	engine := engine.NewEngine()
	if engine == nil {
		t.Fatal("Failed to create new engine")
	}
	if engine.WorldSystem == nil {
		t.Fatal("WorldSystem should not be nil")
	}
	if !engine.Running {
		t.Fatal("Engine should be in a running state")
	}
	if engine.StartTime.IsZero() {
		t.Fatal("StartTime should be set")
	}
	if engine.CurrentTime.IsZero() {
		t.Fatal("CurrentTime should be set")
	}
	println("Engine created successfully with WorldSystem ID:", engine.WorldSystem.ID)
}
