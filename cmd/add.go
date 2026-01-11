package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/FabriOrtiz40/CLI-Task-Manager/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a new task to your TODO list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		dbConn, err := db.Open()
		if err != nil {
			fmt.Println("Error opening task database:", err)
			os.Exit(1)

		}
		defer dbConn.Close()

		err = db.AddTask(dbConn, task)
		if err != nil {
			fmt.Println("Error adding task:", err)
			os.Exit(1)

		}

		fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
