package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/marioharper/commuter/cmd/utils"
	"github.com/marioharper/commuter/directions"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a saved location",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		workReader := bufio.NewReader(os.Stdin)

		// Get location
		fmt.Print("Enter location name: ")
		name, _ := workReader.ReadString('\n')
		name = strings.TrimSpace(name)

		// Get address
		fmt.Print("Enter location address: ")
		address, _ := workReader.ReadString('\n')
		address = strings.TrimSpace(address)

		// Create location
		work := directions.Location{
			Name:    name,
			Address: address,
		}

		// Save location to config
		utils.AddLocation(ConfigFile, work)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
