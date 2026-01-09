package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do [task number]",
	Short: "Mark a task as complete",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("This is a fake \"do\" command for task %s\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
