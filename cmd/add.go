package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a new task to your TODO list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("This is a fake \"add\" command: %s\n", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
