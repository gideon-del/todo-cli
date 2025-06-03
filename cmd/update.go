package cmd

import (
	"cli/todo/todo"
	"errors"
	"strconv"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the todo tile",
	Long:  `Update a todo title. Provide the id and the title`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		todoId, err := strconv.ParseInt(args[0], 10, 32)

		if err != nil {
			return errors.New("invalid todo id")
		}
		selectedTodo, err := todo.GetTodo(int(todoId))

		if err != nil {
			return err
		}
		selectedTodo.Title = args[1]

		return todo.UpdateTodos(selectedTodo)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
