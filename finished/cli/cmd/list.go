package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(list)
}

var list = &cobra.Command{
	Use:   "list",
	Short: "Print Todos",
	Run: func(cmd *cobra.Command, args []string) {
		// iterate over items
	},
}
