package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [task number]",
	Short: "Remove a task from the list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil || index < 1 {
			fmt.Println("Invalid task number!")
			return
		}

		tasks := loadTasks()
		if index > len(tasks) {
			fmt.Println("Task not found!")
			return
		}

		tasks = append(tasks[:index-1], tasks[index:]...)

		saveTasks(tasks)
		fmt.Println("Task removed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

