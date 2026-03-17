package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Note Title:%v\nNote Content:%v\nCreated At:%v\n", note.title, note.content, note.createdAt)
}

func NewNote(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Title And Content Are Required.")
	}

	return Note{
		title:   title,
		content: content,
	}, nil

}
