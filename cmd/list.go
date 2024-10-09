package cmd

import (
	"fmt"
	"tenessine/task-cli/db"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use: "list",
	Short: "List all tasks",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ts, err := db.GetTasksByStatus(args[0])
		if err != nil {
			fmt.Println("Error getting tasks:", err)
			return
		}
		for _, t := range ts {
			fmt.Printf("%-3d | %-20s | %s\n", t.ID, t.Description, t.Status)
		}
	},
}