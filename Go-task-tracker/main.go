package main

import (
	"bufio"
	"fmt"
	"os"
)

func addTask(task string) {
	file, _ := os.OpenFile("tasks.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	file.WriteString(task + "\n")
	fmt.Println("Task added:", task)
}

func viewTasks() {
	file, err := os.Open("tasks.txt")

	if err != nil {
		fmt.Println("No tasks found")
		return
	}

	scanner := bufio.NewScanner(file)

	fmt.Println("Your Tasks:")
	for scanner.Scan() {
		fmt.Println("-", scanner.Text())
	}
}

func main() {

	var choice int

	fmt.Println("1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Print("3. Exit")

	fmt.Scan(&choice)

	if choice == 1 {

		fmt.Println("Enter task:")
		reader := bufio.NewReader(os.Stdin)
		task, _ := reader.ReadString('\n')

		addTask(task)

	} else if choice == 2 {

		viewTasks()

	} else if choice == 3 {

		fmt.Println("Exiting program")
	} else {

		fmt.Println("Invalid option")

	}
}
