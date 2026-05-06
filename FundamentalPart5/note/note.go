package note

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"note_title"`
	Content   string    `json:"note_content"`
	CreatedAt time.Time `json:"note_created_at"`
}

func New(title string, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("title or content is empty")
	}
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func InputDisplay(prompt string) string {
	fmt.Print(prompt + " ")
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	return strings.TrimSpace(data)
}

func (note Note) WriteToJson() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_") + ".json"
	jsonData, err := json.MarshalIndent(fileName, "", " ")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return os.WriteFile(fileName, jsonData, 0644)
}

func ReadFromJson(fileName string) (Note, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return Note{}, err
	}
	var data Note
	err = json.Unmarshal(content, &data)
	if err != nil {
		return Note{}, err
	}
	return data, nil
}
