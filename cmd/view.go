package cmd

import (
	"fmt"

	"github.com/marioharper/commuter/directions"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View your commute on google maps",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		from, err := Config.GetLocationByName(from)
		if err != nil {
			fmt.Printf(err.Error())
			panic(err)
		}
		to, err := Config.GetLocationByName(to)
		if err != nil {
			fmt.Printf(err.Error())
			panic(err)
		}

		commute := directions.Commute{
			From: from,
			To:   to,
		}

		open.Run(commute.GetMapsURL())
	},
}

func init() {
	RootCmd.AddCommand(viewCmd)
}
