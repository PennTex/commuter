package cmd

import (
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View your commute on google maps",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		open.Run(commute.GetMapsURL())
	},
}

func init() {
	RootCmd.AddCommand(viewCmd)
}
