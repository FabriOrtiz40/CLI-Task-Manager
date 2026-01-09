package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is a fake \"list\" command")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
