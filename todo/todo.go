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

	return todos
}
func SaveTodos(todo []Todo) error {
	jsonText, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(todoFile, jsonText, 0644)
}

func GetTodo(id int) (Todo, error) {
	todos := GetTodos()
	for i := range len(todos) {
		todo := todos[i]
		if todo.Id == id {
			return todo, nil
		}
	}

	return Todo{}, fmt.Errorf("Todo with id %v not found", id)
}

func UpdateTodos(updatedTodo Todo) error {
	var updatedTodos []Todo = []Todo{}
	todos := GetTodos()
	for i := range len(todos) {
		todo := todos[i]

		if todo.Id == updatedTodo.Id {
			updatedTodos = append(updatedTodos, updatedTodo)
		} else {
			updatedTodos = append(updatedTodos, todo)
		}
	}

	return SaveTodos(updatedTodos)
}

func DeleteTodo(todoId int) error {
	todos := GetTodos()
	selectedTodo, err := GetTodo(todoId)
	updatedTodos := []Todo{}
	if err != nil {
		return err
	}
	for i := range len(todos) {
		todo := todos[i]

		if todo.Id != selectedTodo.Id {
			updatedTodos = append(updatedTodos, Todo{
				Title:     todo.Title,
				Id:        i + 1,
				Status:    todo.Status,
				CreatedAt: todo.CreatedAt,
				UpdatedAt: todo.UpdatedAt,
			})
		}
	}

	return SaveTodos(updatedTodos)

}

func ChangeTodoStatus(status string, todoId int) error {
	selectedTodo, err := GetTodo(todoId)

	if err != nil {
		return err
	}

	selectedTodo.Status = status

	return UpdateTodos(selectedTodo)
}
