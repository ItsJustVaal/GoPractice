package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(do)
}

var do = &cobra.Command{
	Use:   "do",
	Short: "Completes a Todo",
	Run: func(cmd *cobra.Command, args []string) {
		var done []int
		for _, arg := range args {
			num, _ := strconv.Atoi(arg)
			done = append(done, num)
		}
		fmt.Println(done)
	},
}
