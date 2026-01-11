package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/FabriOrtiz40/CLI-Task-Manager/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do [task number]",
	Short: "Mark a task as complete",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		idx, err := strconv.Atoi(args[0])
		dbConn, err := db.Open()
		if err != nil {
			fmt.Println("Error opening task database:", err)
			os.Exit(1)
		}
		defer dbConn.Close()

		err = db.DeleteTask(dbConn, idx)
		if err != nil {
			fmt.Println("Error completing task:", err)
			os.Exit(1)
		}

		fmt.Printf("You have completed task %d.\n", idx)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
