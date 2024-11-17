package cmd

import (
	"fmt"
	"practicgo/db"
	"practicgo/tools"

	"github.com/spf13/cobra"
)

var (
	title       string
	description string
	status      string
)

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Insert values into Task",
	Run: func(cmd *cobra.Command, args []string) {
		if title == "" {
			fmt.Println("Error: Title empty")
			return
		}
		if description == "" {
			fmt.Println("Error: Description empty")
			return
		}
		if status == "" {
			fmt.Println("Error: Status empty")
			return
		}

		var statusToSql bool

		if status == "y" {
			statusToSql = true
		} else {
			statusToSql = false
		}

		conn := db.SqlConn()
		tools.Post(conn, title, description, statusToSql)
	},
}

func init() {
	rootCmd.AddCommand(postCmd)

	postCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the task")
	postCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the task")
	postCmd.Flags().StringVarP(&status, "status", "s", "", "Status of the task (y/n)")
}
