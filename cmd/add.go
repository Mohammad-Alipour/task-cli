package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var task string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := loadTasks()
		tasks = append(tasks, task)
		saveTasks(tasks)
		fmt.Println("✅ Task added:", task)
	},
}

func init() {
	addCmd.Flags().StringVarP(&task, "task", "t", "", "Task description")
	rootCmd.AddCommand(addCmd)
}

func loadTasks() []string {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		return []string{}
	}
	var tasks []string
	json.Unmarshal(file, &tasks)
	return tasks
}

func saveTasks(tasks []string) {
	data, _ := json.Marshal(tasks)
	os.WriteFile("tasks.json", data, 0644)
}

