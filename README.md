# simple_task

命令行项目管理工具——人下达 task，Agent 完成，人验收。

任务文件随代码一起提交，配合 `git log` 即可追溯每个任务的开发历史和决策过程，让项目文档天然具备时间线。

## 安装

```bash
go build -o simple_task.exe .
```

## 使用

```
simple_task init                  初始化项目（创建 simple_task/ 和 AGENT.md）
simple_task create --title=xxx     创建新任务
simple_task list [--status=xxx]    列出任务，可按状态过滤
simple_task update --id=1 --status=done  更新任务状态
```

### 工作流

1. `simple_task init` 初始化，生成 `simple_task/` 文件夹和把使用方法加入到 `AGENT.md`
2. `simple_task create --title=确定技术栈` 创建任务文件 `task-N-xxx.md`
3. 让 Agent 读取任务文件并完成，Agent 直接编辑 markdown 内容
4. 验收通过后 `simple_task update --id=1 --status=done`

## 规则

- 元信息（frontmatter 中的 `status` / `created_at`）只能通过 CLI 修改
- 任务描述、执行过程等内容可直接编辑 markdown 文件

## 示例任务文件

```markdown
---
created_at: 2026-01-15T10:30:00+08:00
status: todo # done | todo
---

此处为任务描述和执行过程，Agent 可自由编辑。
```
