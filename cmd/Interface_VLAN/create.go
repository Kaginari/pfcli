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
package Interface_VLAN

import (
	"encoding/json"
	"fmt"
	"github.com/Kaginari/pfcli/functions"
	"github.com/Kaginari/pfcli/lib"
	"github.com/Kaginari/pfcli/models"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",

	Run: func(cmd *cobra.Command, args []string) {
		config := lib.GetConfig()
		pfClient := config.Context()
		service  := lib.InterfaceVlanConstruct(pfClient)
		res , err := service.Create(model)
		if err != nil {
			fmt.Println("un error est occurred")
			// TODO FIN BETTER WAY TO HANDLE ERRORS
		}
		jsonRes, _ := json.Marshal(res)
		rest := functions.JsonOutput(jsonRes)

		fmt.Println(rest)
	},
}
 var model models.InterfaceVLAN
func init() {
	pf := CreateCmd.PersistentFlags()
	pf.StringVarP(&model.If, "if", "i", "", models.IvlanIfDesc)
	pf.StringVarP(&model.Tag, "tag", "t", "", models.IvlanTagDesc)
	pf.StringVarP(&model.Pcp, "pcp", "p", "", models.IvlanPcpDesc)
	pf.StringVarP(&model.Descr, "descr", "d", "", models.IvlanDescrDesc)
}


//pfcli InterfaceVLAN create --if 172.5.6.3 --tag 1 --descr newVlan
