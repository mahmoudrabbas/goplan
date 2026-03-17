package main

import (
	"fmt"

	"example.com/note/note"
)

func main() {

	title, content := getNoteDetails()

	note, err := note.NewNote(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	note.Display()

}

func getNoteDetails() (string, string) {
	title := getInput("Note Title: ")
	content := getInput("Note Content: ")
	return title, content
}

func getInput(prompt string) string {
	fmt.Print(prompt)
	var value string
	fmt.Scan(&value)
	return value
}
