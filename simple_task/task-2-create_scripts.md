---
created_at: 2026-08-11T12:15:03+08:00
status: done # done | todo
---

# 存放本项目脚本

直接write编写bat文件到本项目的scrips文件夹（会自己创建）

执行：

`go build -ldflags="-s -w" -o simple_task.exe .`

然后：

`cp simple_task.exe D:/projects/global_exe/`

---

## 实现过程

1. 在项目根目录创建 `scripts/build.bat`，内容为两行命令。
2. 初次使用 `cp`，用户反馈 CMD 不支持，改为 `copy`。
3. `copy` 使用正斜杠路径报错，改为反斜杠 `D:\projects\global_exe\`。
4. 用户验收通过，`simple_task update --id=2 --status=done` 标记完成。
