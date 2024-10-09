package cmd

import (
	"fmt"
	"tenessine/task-cli/db"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Add a new task",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := db.InsertTask(db.Task{
			Description: args[0],
			Status: "todo",
		})
		if err != nil {
			fmt.Printf("Error adding task: %v\n", err)
			return
		}
		fmt.Println("Task added successfully")
	},
}