# 贡献指南

感谢你对 SimEngine 项目的兴趣！我们欢迎各种形式的贡献，包括代码、文档、测试、bug 报告和功能建议。

## 快速开始

1. **Fork 项目** - 点击 GitHub 页面右上角的 "Fork" 按钮
2. **克隆代码** - `git clone https://github.com/your-username/sim_engine.git`
3. **创建分支** - `git checkout -b feature/your-feature`
4. **开发和测试** - 编写代码，确保测试通过
5. **提交更改** - `git commit -m "feat: your feature description"`
6. **推送代码** - `git push origin feature/your-feature`
7. **创建 PR** - 在 GitHub 上创建 Pull Request

## 贡献类型

### 🐛 Bug 报告

发现了 bug？请帮助我们修复它：

1. 检查 [Issues](https://github.com/jelech/sim_engine/issues) 确保问题未被报告
2. 使用 bug 报告模板创建新的 issue
3. 提供详细的重现步骤
4. 包含系统信息和错误日志

### ✨ 功能请求

有好的想法？我们很乐意听到：

1. 检查现有的 [Issues](https://github.com/jelech/sim_engine/issues) 和 [Discussions](https://github.com/jelech/sim_engine/discussions)
2. 创建功能请求 issue，详细描述你的想法
3. 解释为什么这个功能对用户有帮助
4. 如果可能，提供设计草图或伪代码

### 📝 文档改进

文档永远可以更好：

- 修复拼写错误或语法错误
- 改进现有文档的清晰度
- 添加缺失的文档
- 翻译文档到其他语言

### 💻 代码贡献

贡献代码前，请阅读 [开发指南](docs/development.md)。

## 开发环境设置

### 系统要求

- Go 1.20+
- Git 2.0+
- 支持的操作系统：Linux, macOS, Windows

### 安装依赖

```bash
# 克隆项目
git clone https://github.com/jelech/sim_engine.git
cd sim_engine

# 安装 Go 依赖
go mod tidy

# 安装开发工具
make install-tools
```

### 验证设置

```bash
# 运行测试
make test

# 检查代码质量
make lint

# 构建项目
make build
```

## 编码标准

### Go 代码规范

- 遵循 [Go 官方编码规范](https://golang.org/doc/effective_go.html)
- 使用 `gofmt` 和 `goimports` 格式化代码
- 为公开的函数和类型编写文档注释
- 编写有意义的测试

### 提交信息规范

使用 [Conventional Commits](https://www.conventionalcommits.org/) 格式：

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

**类型 (type)：**
- `feat`: 新功能
- `fix`: Bug 修复
- `docs`: 文档更新
- `style`: 代码格式调整（不影响功能）
- `refactor`: 重构（不修复 bug 也不添加功能）
- `perf`: 性能优化
- `test`: 添加或修改测试
- `chore`: 其他维护性更改

**示例：**
```
feat(gym): add support for continuous action spaces

This change introduces support for continuous action spaces in the Gym interface,
allowing for more sophisticated reinforcement learning environments.

Closes #123
```

## Pull Request 流程

### 开发流程

1. **创建功能分支**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **开发和测试**
   - 编写代码
   - 添加或更新测试
   - 确保所有测试通过
   - 更新相关文档

3. **代码检查**
   ```bash
   make lint    # 代码质量检查
   make test    # 运行测试
   make build   # 构建验证
   ```

4. **提交更改**
   ```bash
   git add .
   git commit -m "feat: your feature description"
   ```

5. **保持更新**
   ```bash
   git fetch upstream
   git rebase upstream/main
   ```

6. **推送和创建 PR**
   ```bash
   git push origin feature/your-feature-name
   ```

### PR 要求

你的 Pull Request 应该：

- [ ] 有清晰的标题和描述
- [ ] 关联相关的 issue（如果有）
- [ ] 包含适当的测试
- [ ] 更新相关文档
- [ ] 通过所有 CI 检查
- [ ] 没有合并冲突

### PR 模板

```markdown
## 描述
简要描述这个 PR 的目的和所做的更改。

## 相关 Issue
关闭 #123

## 更改类型
- [ ] Bug 修复
- [ ] 新功能
- [ ] 重构
- [ ] 文档更新
- [ ] 其他（请说明）

## 测试
- [ ] 添加了新的测试
- [ ] 所有测试通过
- [ ] 手动测试通过

## 检查清单
- [ ] 代码遵循项目编码规范
- [ ] 自己审查了代码
- [ ] 添加了适当的注释
- [ ] 更新了相关文档
- [ ] 没有新的编译警告
```

## 测试指南

### 测试类型

- **单元测试**: 测试单个函数或方法
- **集成测试**: 测试组件间的交互
- **基准测试**: 性能测试

### 编写测试

```go
func TestEngine_Start(t *testing.T) {
    engine := NewEngine(DefaultConfig())
    
    err := engine.Start(context.Background())
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    
    if engine.GetStatus() != StatusRunning {
        t.Errorf("Expected status Running, got %v", engine.GetStatus())
    }
}
```

### 运行测试

```bash
# 运行所有测试
make test

# 运行特定包的测试
go test ./pkg/engine

# 运行基准测试
make bench

# 生成覆盖率报告
make coverage
```

## 代码审查

### 审查原则

我们的代码审查关注：

- **正确性**: 代码是否正确实现了功能
- **可读性**: 代码是否易于理解和维护
- **性能**: 是否有明显的性能问题
- **安全性**: 是否存在安全漏洞
- **测试**: 是否有足够的测试覆盖

### 审查流程

1. **自动检查**: CI 会自动运行测试和代码检查
2. **同行审查**: 至少需要一个维护者的批准
3. **修改和改进**: 根据反馈调整代码
4. **最终审查**: 确保所有问题都已解决

## 社区行为准则

### 我们的承诺

为了建设一个开放和包容的环境，我们承诺：

- 使用包容性的语言
- 尊重不同的观点和经验
- 接受建设性的批评
- 关注社区的最佳利益
- 对其他社区成员表示同情

### 不当行为

不可接受的行为包括：

- 使用性别化语言或意象，以及不受欢迎的性关注或示好
- 恶意评论、人身攻击或政治攻击
- 公开或私下的骚扰
- 未经许可发布他人的私人信息
- 其他在职业环境中可能被认为不当的行为

### 举报

如果你遇到或观察到不当行为，请联系项目维护者。所有投诉都会被审查和调查。

## 获取帮助

### 文档资源

- [用户指南](docs/user-guide.md)
- [API 文档](docs/api.md)
- [架构设计](docs/architecture.md)
- [开发指南](docs/development.md)

### 社区支持

- **GitHub Discussions**: 一般问题和讨论
- **GitHub Issues**: Bug 报告和功能请求
- **Email**: jelech@example.com（紧急问题）

### 常见问题

**Q: 我是 Go 新手，可以贡献吗？**
A: 当然可以！我们欢迎所有技能水平的贡献者。从文档改进或简单的 bug 修复开始是很好的选择。

**Q: 如何找到适合初学者的任务？**
A: 查看标记为 "good first issue" 或 "help wanted" 的 issues。

**Q: 我的 PR 被拒绝了怎么办？**
A: 不要气馁！审查意见通常是为了帮助改进代码。根据反馈调整并重新提交。

**Q: 我可以同时处理多个 issue 吗？**
A: 建议一次专注于一个 issue，特别是对于初学者。

## 许可证

通过贡献代码，你同意你的贡献将在与项目相同的 [MIT 许可证](LICENSE) 下发布。

## 致谢

感谢所有为这个项目做出贡献的人！你们的努力让 SimEngine 变得更好。

---

再次感谢你的贡献！如果你有任何问题，请随时在 GitHub 上联系我们。
