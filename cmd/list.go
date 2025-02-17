package cmd

import (
	"fmt"
	"os"
	"encoding/json"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.ReadFile("tasks.json")
		if err != nil {
			fmt.Println("❌ No tasks found!")
			return
		}
		var tasks []string
		json.Unmarshal(file, &tasks)
		fmt.Println("📋 Task List:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

