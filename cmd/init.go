package cmd

import (
	"fmt"

	"github.com/marioharper/commuter/directions"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init your commuter",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		directions.Init()
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
