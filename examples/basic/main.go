package main

import (
	"fmt"
)

func main() {
	fmt.Println("SimEngine Basic Example")
	fmt.Println("=======================")

	/*
		// 创建仿真配置
		config := engine.DefaultConfig()
		config.TimeStep = time.Millisecond * 100
		config.EnableLogging = true

		fmt.Printf("Configuration:\n")
		fmt.Printf("  Time Step: %v\n", config.TimeStep)
		fmt.Printf("  Real Time Mode: %v\n", config.RealTimeMode)
		fmt.Printf("  Max Events: %d\n", config.MaxEvents)
		fmt.Printf("  Logging: %v\n", config.EnableLogging)
		fmt.Println()

		// 导入需要的包
		import (
			"context"
			"log"
			"github.com/jelech/sim_engine/pkg/events"
		)

		// 创建仿真引擎
		sim := engine.NewSimulationEngine(config)

		// 创建一些示例事件
		event1 := events.NewEvent("hello", "Hello from SimEngine!")
		event2 := events.NewEvent("world", "This is a simulation event")
		event3 := events.NewEvent("demo", map[string]interface{}{
			"message": "Demo event with data",
			"value":   42,
			"timestamp": time.Now().Unix(),
		})

		// 调度事件
		now := time.Now()
		sim.ScheduleEvent(event1, now.Add(1*time.Second))
		sim.ScheduleEvent(event2, now.Add(2*time.Second))
		sim.ScheduleEvent(event3, now.Add(3*time.Second))

		fmt.Println("Events scheduled. Starting simulation...")

		// 启动仿真
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := sim.Start(ctx); err != nil {
			log.Fatal("Failed to start simulation:", err)
		}

		// 监控仿真状态
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Simulation timeout reached")
				return
			case <-ticker.C:
				status := sim.GetStatus()
				simTime := sim.GetCurrentTime()
				fmt.Printf("Status: %s, Sim Time: %v\n", status, simTime)

				if status == engine.StatusStopped {
					fmt.Println("Simulation completed")
					return
				}
			}
		}
	*/

	fmt.Println("Basic example completed!")
	fmt.Println("Note: This example will be functional once the core engine is implemented.")
}
