/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"practicgo/db"
	"practicgo/tools"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		conn := db.SqlConn()
		tools.Get(conn)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
