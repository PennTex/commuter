package cmd

import (
	"bufio"
	"os"

	"github.com/PennTex/commuter/cmd/utils"
	"github.com/PennTex/commuter/directions"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a saved location",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		addressValidator := directions.GoogleMapsAddressValidator{}
		workReader := bufio.NewReader(os.Stdin)

		work := directions.Location{
			Name:    utils.GetLocationNameFromUser(workReader),
			Address: utils.GetLocationAddressFromUser(addressValidator, workReader),
		}

		Config.AddLocation(work)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
