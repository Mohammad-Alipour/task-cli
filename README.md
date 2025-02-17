# 📝 Task CLI - Simple Task Manager

A lightweight CLI tool to manage tasks using **Go**. Tasks are stored in a JSON file, no database required.

## 🚀 Features
✅ Add a new task  
✅ List all tasks  
✅ Remove or mark a task as done (coming soon)

## 🛠 Installation & Usage
Clone the repository:
```sh
git clone https://github.com/username/task-cli.git
cd task-cli
go mod tidy

Add a new task:

go run main.go add --task "Learn Golang"

List all tasks:

go run main.go list
