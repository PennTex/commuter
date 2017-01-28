package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/PennTex/commuter/cmd/utils"
	"github.com/PennTex/commuter/directions"
	"github.com/spf13/cobra"
)

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

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init your commuter",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\n\n %s \n\n", logo)

		if i := Config.GetLocations(); len(i) > 0 {
			fmt.Println("Looks like you've already initialized commuter! \nTry `$ commuter list` to view your stored locations.")
			return
		}

		workReader := bufio.NewReader(os.Stdin)
		addressValidator := directions.GoogleMapsAddressValidator{}

		work := directions.Location{
			Name:    "work",
			Address: utils.GetLocationAddressFromUser(addressValidator, workReader),
		}

		Config.AddLocation(work)

		home := directions.Location{
			Name:    "home",
			Address: utils.GetLocationAddressFromUser(addressValidator, workReader),
		}

		Config.AddLocation(home)
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
