package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var todoFile = "todo.json"

type Todo struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(title string, id int, status string) Todo {
	return Todo{
		Id:        id,
		Title:     title,
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

}

func GetTodos() []Todo {
	jsonFile, err := os.ReadFile(todoFile)

	if err != nil {
		fmt.Println(err)
		return []Todo{}
	}

	var todos []Todo

	err = json.Unmarshal(jsonFile, &todos)

	if err != nil {
		fmt.Println(err)
		return []Todo{}
	}
	fmt.Println(todos)
	return todos
}
func SaveTodos(todo []Todo) error {
	jsonText, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(todoFile, jsonText, 0644)
}
