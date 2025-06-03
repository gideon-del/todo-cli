package cmd

import (
	"cli/todo/todo"
	"errors"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a todo by it's id",
	Long:  "Delete a todo by it's id",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		todoId, err := strconv.ParseInt(args[0], 10, 34)
		if err != nil {
			return errors.New("invalid id")
		}

		return todo.DeleteTodo(int(todoId))
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
