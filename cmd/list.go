package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of your current available locations",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		for idx, location := range Locations {
			fmt.Printf("%d) %s - %s\n", idx+1, location.Name, location.Address)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
