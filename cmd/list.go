package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "列出所有任务",
		RunE:  runList,
	}
	cmd.Flags().String("status", "", "按状态过滤: todo | done")
	rootCmd.AddCommand(cmd)
}

func runList(cmd *cobra.Command, args []string) error {
	statusFilter, _ := cmd.Flags().GetString("status")

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("获取当前目录失败: %w", err)
	}

	taskDir := filepath.Join(cwd, "simple_task")
	if info, err := os.Stat(taskDir); err != nil || !info.IsDir() {
		return fmt.Errorf("simple_task 文件夹不存在，请先运行 simple_task init")
	}

	entries, err := os.ReadDir(taskDir)
	if err != nil {
		return fmt.Errorf("读取 simple_task/ 失败: %w", err)
	}

	taskRe := regexp.MustCompile(`^task-(\d+)-(.+)\.md$`)
	printed := 0

	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		m := taskRe.FindStringSubmatch(e.Name())
		if m == nil {
			continue
		}
		id, title := m[1], m[2]

		status := readStatus(filepath.Join(taskDir, e.Name()))

		if statusFilter != "" && status != statusFilter {
			continue
		}

		fmt.Printf("  [%s] #%s  %s\n", status, id, title)
		printed++
	}

	if printed == 0 {
		if statusFilter != "" {
			fmt.Printf("没有状态为 %s 的任务\n", statusFilter)
		} else {
			fmt.Println("暂无任务")
		}
	}

	return nil
}

func readStatus(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "?"
	}

	lines := strings.Split(string(data), "\n")
	inFrontmatter := false

	for _, line := range lines {
		if strings.TrimSpace(line) == "---" {
			if !inFrontmatter {
				inFrontmatter = true
				continue
			} else {
				break
			}
		}
		if inFrontmatter && strings.HasPrefix(strings.TrimSpace(line), "status:") {
			// 提取 status 值，忽略注释
			parts := strings.Fields(strings.TrimPrefix(strings.TrimSpace(line), "status:"))
			if len(parts) > 0 {
				return parts[0]
			}
		}
	}

	return "?"
}
