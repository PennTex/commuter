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

		locationName := utils.GetLocationNameFromUser(workReader)
		work := directions.Location{
			Name:    locationName,
			Address: utils.GetLocationAddressFromUser(locationName, addressValidator, workReader),
		}

		Config.AddLocation(work)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
