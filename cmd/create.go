package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "创建新任务",
		RunE:  runCreate,
	}
	cmd.Flags().String("title", "", "任务标题（必填）")
	cmd.MarkFlagRequired("title")
	rootCmd.AddCommand(cmd)
}

func runCreate(cmd *cobra.Command, args []string) error {
	title, _ := cmd.Flags().GetString("title")

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("获取当前目录失败: %w", err)
	}

	taskDir := filepath.Join(cwd, "simple_task")
	if info, err := os.Stat(taskDir); err != nil || !info.IsDir() {
		return fmt.Errorf("simple_task 文件夹不存在，请先运行 simple_task init")
	}

	// 找到下一个任务编号
	nextID := nextTaskID(taskDir)
	filename := fmt.Sprintf("task-%d-%s.md", nextID, title)
	filePath := filepath.Join(taskDir, filename)

	// 获取 ISO 8601 时间（带时区偏移）
	now := time.Now().Format("2006-01-02T15:04:05-07:00")

	frontmatter := fmt.Sprintf(`---
created_at: %s
status: todo # done | todo
---

xxx任务的具体描述，执行过程和执行结果，验收想法等记录
`, now)

	if err := os.WriteFile(filePath, []byte(frontmatter), 0644); err != nil {
		return fmt.Errorf("创建任务文件失败: %w", err)
	}

	fmt.Printf("已创建任务: %s\n", filename)
	return nil
}

// nextTaskID 扫描 simple_task/ 下所有 task-N-*.md，返回最大 N+1
func nextTaskID(dir string) int {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 1
	}

	re := regexp.MustCompile(`^task-(\d+)-.+\.md$`)
	ids := []int{}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		m := re.FindStringSubmatch(e.Name())
		if m != nil {
			var n int
			fmt.Sscanf(m[1], "%d", &n)
			ids = append(ids, n)
		}
	}

	if len(ids) == 0 {
		return 1
	}

	sort.Ints(ids)
	return ids[len(ids)-1] + 1
}
