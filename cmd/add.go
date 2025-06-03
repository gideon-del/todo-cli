package cmd

import (
	"cli/todo/todo"
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This is used to add todo",
	Long:  `This is used to add a todo`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todoTitle := args[0]
		todos := todo.GetTodos()

		newTodo := todo.New(todoTitle, len(todos)+1, "in-progress")
		updatedTodos := append(todos, newTodo)
		err := todo.SaveTodos(updatedTodos)

		if err != nil {
			fmt.Println(err)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
