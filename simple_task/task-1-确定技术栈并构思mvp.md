---
created_at: 2026-08-11T11:28:00+08:00
status: done # done | todo
---

# 目的

这个项目是一个cli，方便Agent使用来控制项目的task，人负责下达task，agent负责完成，人负责验收

# 库

cli参数解析使用corba：理由是ai推荐并且很火，尝试一下

# 具体任务

最后要得到一个simple_task.exe

使用子命令+flag的格式，下列介绍基本作用的时候顺便介绍格式。

## 基本命令

### init

`simple_task init` or `simple_task init --path=D:/xx/`：在cli运行的路径或指定的path路径下下，先确认是否存在simple_task文件夹如果不存在先创建这个文件夹，如果已经存在那么说明 `D:/xx/xx 已经初始化`，如果不存在就创建一个，返回创建成功

并且再检测路径下是否存在`AGENT.md`文件，如果没有就新建一个添加以下内容，如果有就在已有的文件末尾添加以下内容

```markdown

# simple_task

这是一个命令行项目管理工具，使用 `simple_task help` 获取使用方式

```

### create

`simple_task create --title=xxxx`：在cli运行的路径下下，先确认是否存在simple_task如果不存在则叫先初始化命令，如果存在则创造一个markdown文件

task的文件名称格式是：`task-1-title.md`

task内部使用frontmatter记录元信息，创建task时候直接写入：

时间表示使用ISO 8601 格式当中的时区偏移，先获取本机系统时间然后再得到偏移，注意status后面的#注释内容也要写入方便知道有什么状态

```markdown
---
created_at: 2026-08-11T15:30:00+08:00
status: todo # done | todo
---

```

每个项目下的simple_task文件夹就存放每个任务

格式是`task-xx-名称.md`

### update

`simple_task update --id=1 --status=done`：把id为1的任务的状态改为done，注意task的meta info只能够使用命令行来修改，内容让agent可以直接修改

### agent

`simple_task help`：直接返回一段说明，说明create就是创建，然后update更新状态，以及这个cli的操作逻辑，给以上的例子，元信息只能用cli修改，内容可以直接修改文件

---

# 执行记录

## 实现

使用 cobra（go.mod 中已是 cobra，非 corba）实现，目录结构：

```
cmd/
├── root.go      # 根命令 + 自定义 help（含完整工作流说明和示例）
├── init.go      # simple_task init [--path=...]
├── create.go    # simple_task create --title=xxx
├── update.go    # simple_task update --id=1 --status=done
└── list.go      # simple_task list [--status=xxx]
main.go          # 入口
```

各命令行为：

| 命令 | 行为 |
|---|---|
| `simple_task help` | 显示工作流、规则、所有命令示例 |
| `simple_task init [--path=...]` | 创建 `simple_task/` 目录 + 写入/追加 `AGENT.md`（已存在则提示） |
| `simple_task create --title=xxx` | 扫描已有编号取最大 N+1，创建 `task-N-xxx.md`，写入 frontmatter（ISO 8601 时间 + `status: todo # done | todo`） |
| `simple_task list [--status=xxx]` | 列出所有任务，可按 status 过滤 |
| `simple_task update --id=1 --status=done` | 精确替换 frontmatter 中的 `status:` 行，校验只允许 `todo` / `done` |

## 修 bug

- `readStatus` / `updateStatus` 中 `line == "---"` 无法匹配 Windows `\r\n` 换行，改为 `strings.TrimSpace(line) == "---"`
