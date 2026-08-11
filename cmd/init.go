package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "初始化项目，创建 simple_task/ 文件夹和 AGENT.md",
		RunE:  runInit,
	}
	cmd.Flags().String("path", "", "指定初始化路径（默认当前目录）")
	rootCmd.AddCommand(cmd)
}

func runInit(cmd *cobra.Command, args []string) error {
	path, _ := cmd.Flags().GetString("path")
	if path == "" {
		var err error
		path, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("获取当前目录失败: %w", err)
		}
	}

	path = filepath.Clean(path)
	taskDir := filepath.Join(path, "simple_task")

	// 检查 simple_task 文件夹是否已存在
	if info, err := os.Stat(taskDir); err == nil && info.IsDir() {
		fmt.Printf("%s 已经初始化\n", path)
	} else {
		if err := os.MkdirAll(taskDir, 0755); err != nil {
			return fmt.Errorf("创建 simple_task 文件夹失败: %w", err)
		}
		fmt.Println("创建成功")
	}

	// 处理 AGENT.md
	agentPath := filepath.Join(path, "AGENT.md")
	content := `

# simple_task

这是一个命令行项目管理工具，使用 ` + "`simple_task help`" + ` 获取使用方式
`

	if _, err := os.Stat(agentPath); os.IsNotExist(err) {
		if err := os.WriteFile(agentPath, []byte(content), 0644); err != nil {
			return fmt.Errorf("创建 AGENT.md 失败: %w", err)
		}
	} else {
		f, err := os.OpenFile(agentPath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("打开 AGENT.md 失败: %w", err)
		}
		defer f.Close()
		if _, err := f.WriteString(content); err != nil {
			return fmt.Errorf("追加 AGENT.md 失败: %w", err)
		}
	}

	return nil
}
