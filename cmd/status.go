package cmd

import (
	"cli/todo/todo"
	"errors"
	"strconv"

	"github.com/spf13/cobra"
)

var inProgressCmd = &cobra.Command{
	Use:   "mark-in-progress",
	Args:  cobra.ExactArgs(1),
	Short: "Change a todo's status to in progress",
	Long:  "Change a todo's status to in progress",
	RunE: func(cmd *cobra.Command, args []string) error {
		todoId, err := strconv.ParseInt(args[0], 10, 64)

		if err != nil {
			return errors.New("invalid id format. must be a number")
		}

		return todo.ChangeTodoStatus("in progress", int(todoId))
	},
}
var doneCmd = &cobra.Command{
	Use:   "mark-done",
	Args:  cobra.ExactArgs(1),
	Short: "Change a todo's status to done",
	Long:  "Change a todo's status to done",
	RunE: func(cmd *cobra.Command, args []string) error {
		todoId, err := strconv.ParseInt(args[0], 10, 64)

		if err != nil {
			return errors.New("invalid id format. must be a number")
		}

		return todo.ChangeTodoStatus("done", int(todoId))
	},
}

func init() {
	rootCmd.AddCommand(inProgressCmd, doneCmd)
}
