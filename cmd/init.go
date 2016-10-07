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

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init your commuter",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		var logo = `                                                                                                        
 ██████╗ ██████╗ ███╗   ███╗███╗   ███╗██╗   ██╗████████╗███████╗██████╗ 
██╔════╝██╔═══██╗████╗ ████║████╗ ████║██║   ██║╚══██╔══╝██╔════╝██╔══██╗
██║     ██║   ██║██╔████╔██║██╔████╔██║██║   ██║   ██║   █████╗  ██████╔╝
██║     ██║   ██║██║╚██╔╝██║██║╚██╔╝██║██║   ██║   ██║   ██╔══╝  ██╔══██╗
╚██████╗╚██████╔╝██║ ╚═╝ ██║██║ ╚═╝ ██║╚██████╔╝   ██║   ███████╗██║  ██║
 ╚═════╝ ╚═════╝ ╚═╝     ╚═╝╚═╝     ╚═╝ ╚═════╝    ╚═╝   ╚══════╝╚═╝  ╚═╝
                                                                         
                                   _._
                              _.-="_-         _
                         _.-="   _-          | ||"""""""---._______     __..
             ___.===""""-.______-,,,,,,,,,,,,'-''----" """""       """""  __'
      __.--""     __        ,'                   o \           __        [__|
 __-""=======.--""  ""--.=================================.--""  ""--.=======:
]       [w] : /        \ : |========================|    : /        \ :  [w] :
V___________:|          |: |========================|    :|          |:   _-"
 V__________: \        / :_|=======================/_____: \        / :__-"
 -----------'  ""____""  '-------------------------------'  ""____""                                                                                                                                               
		`

		fmt.Printf("\n\n %s \n\n", logo)

		// Get work address
		workReader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter work address: ")
		workAddress, _ := workReader.ReadString('\n')
		workAddress = strings.TrimSpace(workAddress)
		work := directions.Location{
			Name:    "work",
			Address: workAddress,
		}
		utils.AddLocation(ConfigFile, work)

		// Get home address
		homeReader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter home address: ")
		homeAddress, _ := homeReader.ReadString('\n')
		homeAddress = strings.TrimSpace(homeAddress)
		home := directions.Location{
			Name:    "home",
			Address: homeAddress,
		}
		utils.AddLocation(ConfigFile, home)
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
