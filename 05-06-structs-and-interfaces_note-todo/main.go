package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/phubadeepjs/note/note"
	"github.com/phubadeepjs/note/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello!")

	title, content := getNoteData()

	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	printSomething(todo)

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func printSomething(value interface{}) {
	intVal, ok := value.(int)
	if !ok {
		fmt.Println("Integer:", intVal)
		return
	}

	folatVal, ok := value.(float64)
	if !ok {
		fmt.Println("Float64:", folatVal)
		return
	}

	stringVal, ok := value.(string)
	if !ok {
		fmt.Println("String:", stringVal)
		return
	}
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float64:", value)
	// case string:
	// 	fmt.Println("String:", value)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Title:")
	content := getUserInput("Content:")

	return title, content
}
