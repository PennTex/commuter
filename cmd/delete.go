// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/marioharper/commuter/cmd/utils"
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

		utils.DeleteLocation(ConfigFile, name)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
