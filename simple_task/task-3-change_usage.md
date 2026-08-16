---
created_at: 2026-08-16T10:29:49+08:00
status: todo # done | todo
---

> 人下达task，你执行task并记录执行过程，人验收后记录验收看到的东西判断是否完成，你继续执行task，直到人认为结束了再手动修改status

# 目的

现在要大改一下这个命令行工具的用法。

1. 现在这个命令行工具有一个 `simple_task update` 的命令，但是没必要因为这个工具的使用方法是 人下达task，agent执行task并记录执行过程，人验收后记录验收看到的东西判断是否完成，agent继续执行task，直到人认为结束了再手动修改status所以不用这个命令了
2. 现在这个命令行要 init 创建 `AGENTS.md` 然后 `simple_task help` 才知道用法，这个其实没有必要，直接把理念写到create 创建的任务文件中就好，就是只保留 `simple_task create --title=xx` 这个命令，如果不存在 simple_task文件夹就顺便创造，存在的话就正常序号增加创造
3. `simple_task create` 命令创造的文件在 frontmatter下方新加入 `> 人下达task，你执行task并记录执行过程，人验收后记录验收看到的东西判断是否完成，你继续执行task，直到人认为结束了再手动修改status`
4. `simple_task help` 展示的文字也修改一下，不要那些说只能用命令行修改status的话了

## 执行记录

- 删除 `simple_task update`：状态由人手动改 markdown，无需 CLI。
- 删除 `simple_task init` 与 AGENT.md 生成逻辑。
- `simple_task create --title=xx`：`simple_task/` 不存在时自动创建，存在时按最大编号 +1 递增。
- `create` 生成的文件在 frontmatter 下方新增理念提示行。
- `simple_task help` 文案更新，去掉"只能用命令行修改 status"的规则。
- `simple_task list` 保留（按用户反馈恢复），`simple_task/` 不存在时输出"暂无任务"。
- `go build ./...` 编译通过。

# 验收

list和create命令正常运行