package cmd

import (
	"fmt"
	"practicgo/db"
	"practicgo/tools"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	id        int
	statusdel string
)

var putCmd = &cobra.Command{
	Use:   "put",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if id == 0 {
			fmt.Println("Id is empty")
			return
		}

		status, err := strconv.ParseBool(statusdel)
		if err != nil {
			fmt.Println("Status is empty")
			return
		}

		conn := db.SqlConn()
		tools.Put(conn, id, status)
	},
}

func init() {
	rootCmd.AddCommand(putCmd)

	putCmd.Flags().IntVarP(&id, "id", "i", 0, "Id of the column")
	putCmd.Flags().StringVarP(&statusdel, "status", "s", "", "Status of the task true/false")
}
