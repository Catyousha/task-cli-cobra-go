package cmd

import (
	"fmt"
	"strconv"
	"tenessine/task-cli/db"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(markDoneCmd)
}

var markDoneCmd = &cobra.Command{
	Use:   "mark-done",
	Short: "Mark a task as done",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid id:", err)
			return
		}
		_, err = db.UpdateTaskStatus(id, "done")
		if err != nil {
			fmt.Println("Error marking task as done:", err)
			return
		}
		fmt.Println("Task marked as done")
	},
}
