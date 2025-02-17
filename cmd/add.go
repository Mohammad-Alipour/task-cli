package cmd

import (
	"fmt"
	"os"
	"encoding/json"
)

const taskFile = "tasks.json"


type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}


func AddTask(taskName string) {
	tasks := loadTasks()

	
	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}


	newTask := Task{ID: newID, Name: taskName}
	tasks = append(tasks, newTask)

	
	saveTasks(tasks)
	fmt.Println("Task added:", taskName)
}


func loadTasks() []Task {
	file, err := os.ReadFile(taskFile)
	if err != nil {
		return []Task{}
	}

	var tasks []Task
	json.Unmarshal(file, &tasks)
	return tasks
}

// ذخیره تسک‌ها در فایل
func saveTasks(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	os.WriteFile(taskFile, data, 0644)
}

