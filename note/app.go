package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputable interface {
	saver
	Display()
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	return data.Save()
}

func printAnything(v interface{}) {
	fmt.Println(v)
}

func main() {

	// printAnything(1)

	title, content := getNoteDetails()

	text := getTodoText()

	note, err := note.NewNote(title, content)

	todo, err := todo.NewTodo(text)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(note)

	outputData(todo)

	// note.Display()
	// todo.Display()

}

func getNoteDetails() (string, string) {
	title := getInput("Note Title: ")
	content := getInput("Note Content: ")
	return title, content
}

func getTodoText() string {
	text := getInput("Todo Text:")
	return text
}

func getInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
