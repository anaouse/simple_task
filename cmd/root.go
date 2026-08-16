package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "simple_task",
	Short: "命令行项目管理工具，方便 Agent 控制项目 task",
	Long: `simple_task —— 人下达 task，Agent 完成，人验收。

工作流：
  1. simple_task create --title=xxx    创建新任务（不存在 simple_task/ 时会自动创建，生成 task-N-xxx.md）
  2. simple_task list [--status=xxx]   列出任务，可按状态过滤
  3. 让 Agent 读取任务文件并完成，Agent 可直接编辑 markdown 文件内容
  4. 人验收通过后，手动把任务文件 frontmatter 中的 status 改为 done

示例：
  simple_task create --title=确定技术栈并构思mvp
  simple_task list
  simple_task list --status=todo`,
}

func Execute() {
	rootCmd.Execute()
}
