package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "task-cli",
	Short: "Task is a CLI task manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Task CLI")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}