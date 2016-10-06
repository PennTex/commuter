package cmd

import (
	"bufio"
	"os"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init your commuter",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		f, err := os.Create("config.json")
    	check(err)

		workReader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter work address: ")
		workAddress, _ := workReader.ReadString('\n')
		workAddress = strings.TrimSpace(workAddress)

		homeReader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter home address: ")
		homeAddress, _ := homeReader.ReadString('\n')
		homeAddress = strings.TrimSpace(homeAddress)

		w := bufio.NewWriter(f)
		w.WriteString("[ \n \t { \n \t \t \"name\" : \"work\" , \n \t \t \"address\": \"" + workAddress + "\" \n \t }, \n \t { \n \t \t \"name\" : \"home\" , \n \t \t \"address\": \"" + homeAddress + "\" \n \t } \n ]")
		w.Flush()
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}