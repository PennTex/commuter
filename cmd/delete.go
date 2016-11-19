package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a saved location",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		workReader := bufio.NewReader(os.Stdin)

		// Get location name
		fmt.Print("Enter location name to delete: ")
		name, _ := workReader.ReadString('\n')
		name = strings.TrimSpace(name)

		Config.DeleteLocation(name)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
