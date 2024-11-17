package cmd

import (
	"fmt"
	"practicgo/db"
	"practicgo/tools"

	"github.com/spf13/cobra"
)

var idtask int

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		if idtask == 0 {
			fmt.Println("Id is empty")
			return
		}

		conn := db.SqlConn()
		tools.Delete(conn, idtask)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntVarP(&idtask, "id", "i", 0, "Id column for delete")
}
