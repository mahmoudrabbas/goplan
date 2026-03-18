package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("Note Title:%v\nNote Content:%v\nCreated At:%v\n", note.Title, note.Content, note.CreatedAt)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func NewNote(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Title And Content Are Required.")
	}

	return Note{
		Title:   title,
		Content: content,
	}, nil

}
