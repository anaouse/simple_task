package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "simple_task",
	Short: "命令行项目管理工具，方便 Agent 控制项目 task",
	Long: `simple_task —— 人下达 task，Agent 完成，人验收。

工作流：
  1. simple_task init                 初始化项目（创建 simple_task/ 文件夹和 AGENT.md）
  2. simple_task create --title=xxx    创建一个新任务（在 simple_task/ 下生成 task-N-xxx.md）
  3. simple_task list [--status=xxx]   列出任务，可按状态过滤
  4. 让 Agent 读取任务文件并完成，Agent 可直接编辑文件内容
  5. simple_task update --id=1 --status=done  验收通过后更新任务状态

规则：
  - 元信息（frontmatter 中的 status / created_at）只能用 CLI 修改
  - 任务描述、执行过程等内容可以直接编辑 markdown 文件，在用户确保满意叫你用命令行把把状态改为 done 后记得记录好过程

示例：
  simple_task init
  simple_task init --path=D:/my_project/
  simple_task create --title=确定技术栈并构思mvp
  simple_task list
  simple_task list --status=todo
  simple_task update --id=1 --status=done`,
}

func Execute() {
	rootCmd.Execute()
}
