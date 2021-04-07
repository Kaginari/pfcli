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
package firewallRule

import (
	"encoding/json"
	"fmt"
	"github.com/Kaginari/pfcli/functions"
	"github.com/Kaginari/pfcli/lib"
	"github.com/Kaginari/pfcli/models"
	"github.com/spf13/cobra"
)

// DeleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Run: func(cmd *cobra.Command, args []string) {
		config := lib.GetConfig()
		pfClient:=config.Context()
		service:=lib.FirewallRuleServiceConstruct(pfClient)
		res,err:=service.Delete(DeleteModel)
		if err != nil {
			fmt.Println("un error est occurred")
			// TODO FIN BETTER WAY TO HANDLE ERRORS
		}
		jsonRes, _ := json.Marshal(res)
		rest := functions.JsonOutput(jsonRes)

		fmt.Println(rest)
	},
}


var DeleteModel models.DeleteFirewallRule
func init() {
	pf := DeleteCmd.PersistentFlags()
	pf.StringVarP(&DeleteModel.Tracker, "tracker", "t", "", "Specify the rule tracker ID to delete")
	pf.BoolVar(&DeleteModel.Apply, "apply",true, models.FWRApplyDesc)
}
//pfcli firewallRule delete --tracker 1616663222 --apply