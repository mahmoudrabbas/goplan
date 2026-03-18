package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	err = note.Save()
	if err != nil {
		fmt.Println(err, "Sorry, Failed To Save Note.")
	}

}

func getNoteDetails() (string, string) {
	title := getInput("Note Title: ")
	content := getInput("Note Content: ")
	return title, content
}

func getInput(prompt string) string {
	fmt.Print(prompt)
	// var value string

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	// fmt.Scan(&value)
	return text
}
