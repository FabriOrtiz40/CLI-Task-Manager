package cmd

import (
	"fmt"
	"os"

	"github.com/FabriOrtiz40/CLI-Task-Manager/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		dbConn, err := db.Open()
		if err != nil {
			fmt.Println("Error opening task database:", err)
			os.Exit(1)

		}
		defer dbConn.Close()

		tasks, err := db.ListTasks(dbConn)
		if err != nil {
			fmt.Println("Error listing tasks:", err)
			os.Exit(1)

		}

		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i, task)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
