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
package InterfaceVLAN

import (
	"encoding/json"
	"fmt"
	"github.com/Kaginari/pfcli/functions"
	"github.com/Kaginari/pfcli/lib"
	"github.com/Kaginari/pfcli/models"
	"github.com/spf13/cobra"
)

// DeleteCmd represents the Delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",

	Run: func(cmd *cobra.Command, args []string) {
		config := lib.GetConfig()
		pfClient:=config.Context()
		service:=lib.InterfaceVlanConstruct(pfClient)
		res,err:=service.Delete(modelDelete)
		if err != nil {
			fmt.Println("un error est occurred")
			// TODO FIN BETTER WAY TO HANDLE ERRORS
		}
		jsonRes, _ := json.Marshal(res)
		rest := functions.JsonOutput(jsonRes)

		fmt.Println(rest)
	},
}
var modelDelete models.InterfaceVlanDelete
func init() {
	pf := DeleteCmd.PersistentFlags()
	pf.StringVarP(&modelDelete.Vlanif, "vlanif", "v", "", models.IvlanVlanifDesc)
	pf.StringVarP(&modelDelete.Id, "id", "i", "", models.IvlanIdfDesc)
}

//pfcli InterfaceVLAN delete --vlanif vmp1