package cmd

import (
	"fmt"
	"strconv"
	"tenessine/task-cli/db"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(markInProgressCmd)
}

var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "Mark a task as in progress",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid id:", err)
			return
		}
		_, err = db.UpdateTaskStatus(id, "in-progress")
		if err != nil {
			fmt.Println("Error marking task as in progress:", err)
			return
		}
		fmt.Println("Task marked as in progress")
	},
}
