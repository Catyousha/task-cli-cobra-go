package cmd

import (
	"fmt"
	"strconv"
	"tenessine/task-cli/db"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid id:", err)
			return
		}
		_, err = db.UpdateTaskDescription(id, args[1])
		if err != nil {
			fmt.Println("Error updating task:", err)
			return
		}
		fmt.Println("Task updated successfully")
	},
}
