# simple_task

命令行项目管理工具——人下达 task，Agent 完成，人验收。

任务文件随代码一起提交，配合 `git log` 即可追溯每个任务的开发历史和决策过程，让项目文档天然具备时间线。

## 安装

```bash
go build -o simple_task.exe .
```

## 使用

```
simple_task create --title=xxx   创建新任务（不存在 simple_task/ 时会自动创建，生成 task-N-xxx.md）
simple_task list [--status=xxx]  列出任务，可按状态过滤
```

### 工作流

1. `simple_task create --title=确定技术栈` 创建任务文件 `task-N-xxx.md`
2. `simple_task list` 列出任务
3. 让 Agent 读取任务文件并完成，Agent 直接编辑 markdown 内容
4. 验收通过后，人手动把任务文件 frontmatter 中的 `status` 改为 `done`

## 示例任务文件

```markdown
---
created_at: 2026-01-15T10:30:00+08:00
status: todo # done | todo
---

> 人下达task，你执行task并记录执行过程，人验收后记录验收看到的东西判断是否完成，你继续执行task，直到人认为结束了再手动修改status
```
