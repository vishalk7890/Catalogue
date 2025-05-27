package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Title string `json:title`
	Done  bool   `json:done`
}

var tasks []Task

const fileName = "tasks.json"

func main() {
	loadTasks()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nTodo CLI")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")

		fmt.Print("Choose an option: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			addTask(reader)
		case "2":
			listTasks()
		case "3":
			markTaskDone(reader)
		case "4":
			deleteTask(reader)
		case "5":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Invalid option.")
		}

	}
}

func loadTasks() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {

			tasks = []Task{}
			return
		}
		fmt.Println("Error reading tasks:", err)
		return
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error parsing tasks:", err)
		return
	}
}

func addTask(reader *bufio.Reader) {
	fmt.Println("Enter the task title:")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	tasks = append(tasks, Task{Title: title, Done: false})
	saveTasks()
	fmt.Println("Task added:", title)

}

func saveTasks() {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("Error writing tasks to file:", err)
		return

	}
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}
	for i, task := range tasks {
		status := "[ ]"
		if task.Done {
			status = "[x]"

		}
		fmt.Println("%d. %s %s\n", i+1, status, task.Title)
	}
}

func markTaskDone(reader *bufio.Reader) {
	listTasks()
	fmt.Println("Enter the task number to mark as done:")
	input, _ := reader.ReadString('\n')
	index, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || index < 1 || index > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}
	tasks[index-1].Done = true
	saveTasks()
	fmt.Println("Task marked as done.")
}

func deleteTask(reader *bufio.Reader) {
	listTasks()
	fmt.Println("Enter the number of task to delete")
	input, _ := reader.ReadString('\n')
	index, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || index < 1 || index > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}
	tasks = append(tasks[:index-1], tasks[index:]...)
	fmt.Println("Task deleted.")
	saveTasks()
}
