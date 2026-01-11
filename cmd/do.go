package cmd

import (
	"fmt"
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
		if err != nil {
			panic(err)
		}

		dbConn, err := db.Open()
		if err != nil {
			panic(err)
		}
		defer dbConn.Close()

		err = db.DeleteTask(dbConn, idx)
		if err != nil {
			panic(err)
		}

		fmt.Printf("You have completed task %d.\n", idx)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
