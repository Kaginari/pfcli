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
package firewall_rule

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.com/nt-factory/2021/admin/pfcli/functions"
	"gitlab.com/nt-factory/2021/admin/pfcli/models"
	"log"
	"net/http"
)

// DeleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Run: func(cmd *cobra.Command, args []string) {
		DeletRule()
	},
}


var DeleteModel models.DeleteFirewallRule
func init() {
	pf := DeleteCmd.PersistentFlags()
	pf.StringVarP(&DeleteModel.Tracker, "tracker", "t", "", "Specify the rule tracker ID to delete")
	pf.BoolVar(&DeleteModel.Apply, "apply",false, models.FWRDisabledDesc)
}
func DeletRule()  {
	jsonReq, _ := json.Marshal(DeleteModel)
	res := functions.JsonOutput(jsonReq)
	fmt.Println(res)
	req, err := http.NewRequest("DELETE", functions.ViperReadConfig().UrlPfsense+"v1/firewall/rule", bytes.NewBuffer(jsonReq))
	req.Header.Add("Authorization", functions.ViperReadConfig().ClientId + " "+functions.ViperReadConfig().ClientToken)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)


	}

	fmt.Println("response Status : ", resp.Status)
	jsonReq2, _ := json.Marshal(resp.Body)
	res2 := functions.JsonOutput(jsonReq2)
	fmt.Println(res2)
	fmt.Println("response Headers : ", resp.Header)

}