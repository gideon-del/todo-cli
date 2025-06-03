package cmd

import (
	"cli/todo/todo"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List out all todos or filter by status",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			filter := args[0]
			fmt.Println(todo.FilterTodoByStatus(filter))

		} else {
			fmt.Println(todo.GetTodos())
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
