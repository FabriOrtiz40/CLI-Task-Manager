package cmd

import (
	"fmt"

	"github.com/FabriOrtiz40/CLI-Task-Manager/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		dbConn, err := db.Open()
		if err != nil {
			panic(err)
		}
		defer dbConn.Close()

		tasks, err := db.ListTasks(dbConn)
		if err != nil {
			panic(err)
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
