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
		Use:   "update",
		Short: "更新任务状态",
		RunE:  runUpdate,
	}
	cmd.Flags().Int("id", 0, "任务 ID（必填）")
	cmd.Flags().String("status", "", "新状态: todo | done（必填）")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("status")
	rootCmd.AddCommand(cmd)
}

func runUpdate(cmd *cobra.Command, args []string) error {
	id, _ := cmd.Flags().GetInt("id")
	status, _ := cmd.Flags().GetString("status")

	// 校验 status
	validStatus := map[string]bool{"todo": true, "done": true}
	if !validStatus[status] {
		return fmt.Errorf("无效的状态: %s，可选: todo | done", status)
	}

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("获取当前目录失败: %w", err)
	}

	taskDir := filepath.Join(cwd, "simple_task")
	if info, err := os.Stat(taskDir); err != nil || !info.IsDir() {
		return fmt.Errorf("simple_task 文件夹不存在，请先运行 simple_task init")
	}

	// 找到匹配 task-{id}-*.md 的文件
	filePath, err := findTaskFile(taskDir, id)
	if err != nil {
		return err
	}

	// 读取并更新 frontmatter 中的 status
	if err := updateStatus(filePath, status); err != nil {
		return err
	}

	fmt.Printf("任务 %d 状态已更新为: %s\n", id, status)
	return nil
}

func findTaskFile(dir string, id int) (string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", fmt.Errorf("读取 simple_task/ 失败: %w", err)
	}

	pattern := regexp.MustCompile(fmt.Sprintf(`^task-%d-.+\.md$`, id))
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		if pattern.MatchString(e.Name()) {
			return filepath.Join(dir, e.Name()), nil
		}
	}

	return "", fmt.Errorf("未找到 id 为 %d 的任务文件", id)
}

func updateStatus(filePath string, newStatus string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("读取任务文件失败: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	inFrontmatter := false
	updated := false

	for i, line := range lines {
		if strings.TrimSpace(line) == "---" {
			if !inFrontmatter {
				inFrontmatter = true
				continue
			} else {
				// 到达 frontmatter 结束标记，后面的不再处理
				break
			}
		}
		if inFrontmatter && strings.HasPrefix(strings.TrimSpace(line), "status:") {
			lines[i] = fmt.Sprintf("status: %s # done | todo", newStatus)
			updated = true
			break
		}
	}

	if !updated {
		return fmt.Errorf("未在文件中找到 status 字段")
	}

	output := strings.Join(lines, "\n")
	if err := os.WriteFile(filePath, []byte(output), 0644); err != nil {
		return fmt.Errorf("写入任务文件失败: %w", err)
	}

	return nil
}
