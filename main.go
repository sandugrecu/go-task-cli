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
	Content string
	IsDone  bool
}

func NewTask(Content string) Task {
	return Task{
		Content: Content,
		IsDone:  false,
	}
}

func main() {
	//tasks := make([]Task, 0)
	tasks := loadTasks("data.json")
	reader := bufio.NewReader(os.Stdin)

	for {
		displayMenu()
		fmt.Print("Your choise: ")
		text, _ := reader.ReadString('\n')
		chosenOption := strings.TrimSpace(text)
		switch chosenOption {
		case "1":
			tasks = addTask(tasks, reader)

		case "2":
			listTasks(tasks)

		case "3":
			fmt.Print("What is the index of the task you want to change the status: ")
			index, err := getIndex(reader)
			if err != nil {
				fmt.Println("Invalid index. Index must be an integer number.")
				continue
			}
			changeStatus(&tasks, index)

		case "4":
			fmt.Print("What is the index of the task you want to delete: ")
			index, err := getIndex(reader)
			if err != nil {
				fmt.Println("Invalid index. Index must be an integer number.")
				continue
			}
			deleteTask(&tasks, index)

		case "5":
			saveTasks("data.json", tasks)
			fmt.Println("Exiting...")
			os.Exit(0)

		default:
			fmt.Println("Invalid option! Choose a number in range 1-4.")
		}
	}
}

func displayMenu() {
	fmt.Println(`
1. Add task
2. List all tasks 
3. Mark task as completed (By ID)
4. Delete task (By ID)
5. Save tasks and exit`)
}

func addTask(tasks []Task, reader *bufio.Reader) []Task {
	fmt.Println("What is the task:")
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)
	if task == "" {
		// If task is empty the original tasks slice will be returned
		return tasks
	}

	return append(tasks, NewTask(task))
}

func listTasks(tasks []Task) {
	fmt.Println("________________________\nAll available tasks:")
	for index, task := range tasks {
		status := "incomplete"
		if task.IsDone {
			status = "complete"
		}

		fmt.Printf("\nTask %d\nTask: %s. \nStatus %s\n", index, task.Content, status)
	}
	fmt.Println("________________________")
}

func changeStatus(tasks *[]Task, index int) {
	if index < 0 || index >= len(*tasks) {
		fmt.Println("Task index does not exist.")
		return
	}

	(*tasks)[index].IsDone = !(*tasks)[index].IsDone
}

func deleteTask(tasks *[]Task, index int) {
	if index < 0 || index >= len(*tasks) {
		fmt.Println("Task index does not exist.")
		return
	}

	*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)
}

func getIndex(reader *bufio.Reader) (int, error) {
	text, _ := reader.ReadString('\n')
	indexStr := strings.TrimSpace(text)
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		return 0, err
	}
	return index, nil
}

func loadTasks(filename string) []Task {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Could not read file:", err)
		return []Task{}
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Invalid JSON:", err)
		return []Task{}
	}
	return tasks
}

func saveTasks(filename string, tasks []Task) {
	data, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Could not encode json:", err)
		return
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Could not write to file:", err)
		return
	}
}
