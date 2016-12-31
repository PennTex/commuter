package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/marioharper/commuter/directions"
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

func getAddressLocationFromUser(location *directions.Location, reader *bufio.Reader) {
	location.Address = ""

	for location.Address == "" {
		fmt.Printf("Enter %s address: ", location.Name)
		location.Address, _ = reader.ReadString('\n')
		location.Address = strings.TrimSpace(location.Address)

		if location.Address == "" {
			fmt.Printf("Please provide a %s address. \n", location.Name)
		}
	}
}

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

		work := directions.Location{
			Name: "work",
		}
		getAddressLocationFromUser(&work, workReader)
		Config.AddLocation(work)

		home := directions.Location{
			Name: "home",
		}
		getAddressLocationFromUser(&home, workReader)
		Config.AddLocation(home)
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
