/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"gitlab.com/nt-factory/2021/admin/pfcli/cmd/Interface_VLAN"

	"github.com/spf13/cobra"
)

// InterfaceVLANCmd represents the InterfaceVLAN command
var InterfaceVLANCmd = &cobra.Command{
	Use:   "InterfaceVLAN",
	Short: "VLAN interface.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("InterfaceVLAN called")
	},
}

func init() {
}
func InterfaceVLANAddCommands(){
	InterfaceVLANCmd.AddCommand(Interface_VLAN.CreateCmd)
	InterfaceVLANCmd.AddCommand(Interface_VLAN.DeleteCmd)
	InterfaceVLANCmd.AddCommand(Interface_VLAN.ListCmd)


}