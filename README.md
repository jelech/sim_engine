# SimEngine - 事件驱动仿真引擎

[![Go Version](https://img.shields.io/badge/Go-1.20+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Documentation](https://img.shields.io/badge/docs-godoc-blue.svg)](https://godoc.org/github.com/jelech/sim_engine)

SimEngine 是一个高性能的事件驱动仿真引擎，专为强化学习应用设计。它提供了完整的 OpenAI Gym 兼容接口，支持离散和连续动作空间，适用于各种强化学习算法的研究和开发。

## ✨ 特性

- **🎯 事件驱动架构**: 高效的事件调度和处理系统
- **🤖 强化学习支持**: 完整的 OpenAI Gym 兼容接口
- **⚡ 高性能**: 优化的调度器和内存管理
- **🔧 模块化设计**: 可扩展的插件式架构
- **📊 丰富的工具**: 内置统计、可视化和调试工具
- **🌐 多环境支持**: 支持单任务和多任务仿真
- **📈 实时监控**: 仿真状态和性能监控

## 🚀 快速开始

### 安装

```bash
go get github.com/jelech/sim_engine
```

### 基本使用

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/jelech/sim_engine/pkg/engine"
    "github.com/jelech/sim_engine/pkg/gym"
)

func main() {
    // 创建仿真引擎
    config := engine.DefaultConfig()
    engine := engine.NewSimulationEngine(config)
    
    // 创建强化学习环境
    env, err := gym.NewEnvironment("CartPole-v1", nil)
    if err != nil {
        log.Fatal(err)
    }
    
    // 启动仿真
    ctx := context.Background()
    if err := engine.Start(ctx); err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("仿真引擎启动成功！")
}
```

## 📖 文档

- [用户指南](docs/user-guide.md)
- [API 文档](docs/api.md)
- [架构设计](docs/architecture.md)
- [示例教程](docs/examples.md)
- [开发指南](docs/development.md)

## 🏗️ 架构

### 核心组件

- **Engine**: 仿真引擎核心，管理仿真生命周期
- **Events**: 事件定义和处理机制
- **Scheduler**: 高效的事件调度系统
- **Gym**: OpenAI Gym 兼容的强化学习接口
- **RL**: 强化学习算法和工具

## 🎮 示例

### 1. 简单的事件仿真

```go
// 创建事件
event := events.NewEvent("move", map[string]interface{}{
    "entity_id": "player1",
    "position":  []float64{10.0, 20.0},
})

// 调度事件
scheduler.Schedule(event, time.Now().Add(time.Second))
```

更多示例请查看 [examples](examples/) 目录。

## 🧪 测试

运行所有测试：

```bash
go test ./...
```

运行特定包的测试：

```bash
go test ./pkg/engine
go test ./pkg/gym
```

运行基准测试：

```bash
go test -bench=. ./...
```

## 🤝 贡献

我们欢迎所有形式的贡献！请查看 [CONTRIBUTING.md](CONTRIBUTING.md) 了解如何参与项目开发。

### 开发流程

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 📋 路线图

- [ ] 基础事件系统
- [ ] 仿真引擎核心
- [ ] Gym 接口实现
- [ ] 常用强化学习算法
- [ ] 可视化工具
- [ ] 分布式仿真支持
- [ ] Web 监控面板
- [ ] 更多环境模板

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🙏 致谢

感谢以下项目的启发：

- [OpenAI Gym](https://github.com/openai/gym)
- [DeepMind Lab](https://github.com/deepmind/lab)
- [Unity ML-Agents](https://github.com/Unity-Technologies/ml-agents)

## 📞 联系

- 项目主页: [https://github.com/jelech/sim_engine](https://github.com/jelech/sim_engine)
- 问题反馈: [Issues](https://github.com/jelech/sim_engine/issues)
- 讨论交流: [Discussions](https://github.com/jelech/sim_engine/discussions)

---

⭐ 如果这个项目对你有帮助，请给我们一个星标！