package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "practicgo",
	Short: "CLI TODO list with pgsql",
	Long:  `Using this CLI for managment your TODO list. You can PUT, GET, POST, DELETE data from BD and udate status task`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is CLI app - using flag --help to give more information")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
