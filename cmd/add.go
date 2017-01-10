package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/PennTex/commuter/directions"
	"github.com/spf13/cobra"
)

func getLocationName(reader *bufio.Reader) string {
	name := ""

	for name == "" {
		fmt.Print("Enter location name: ")
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if name == "" {
			fmt.Println("Please supply a value for the location's name.")
		}
	}

	return name
}

func getLocationAddress(reader *bufio.Reader) string {
	address := ""

	for address == "" {
		fmt.Print("Enter location address: ")
		address, _ = reader.ReadString('\n')
		address = strings.TrimSpace(address)

		validAddress := directions.AddressIsValid(address)
		if !validAddress {
			address = ""
		}

		if address == "" {
			fmt.Println("Please supply a valid address for the location.")
		}
	}

	return address
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a saved location",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		workReader := bufio.NewReader(os.Stdin)

		work := directions.Location{
			Name:    getLocationName(workReader),
			Address: getLocationAddress(workReader),
		}

		Config.AddLocation(work)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
