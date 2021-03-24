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
package Virtual_IPs

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

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Run: func(cmd *cobra.Command, args []string) {
		createVirtualIPS()
	},
}

func init() {
	createFlags()
}

var VirtualIPS models.Virtual_IPS
func createFlags()  {
	pf := CreateCmd.PersistentFlags()
	pf.StringVarP(&VirtualIPS.Mode, "mode", "m","",models.VertualIPS_Mode_Desc )
	pf.StringVarP(&VirtualIPS.Interface, "interface", "i","",models.VertualIPS_Interface_Desc )
	pf.StringVarP(&VirtualIPS.Subnet, "subnet", "s","",models.VertualIPS_Subnet_Desc )
	pf.StringVarP(&VirtualIPS.Descr, "descr", "d","",models.VertualIPS_Mode_Desc )
	pf.BoolVarP(&VirtualIPS.Noexpand, "noexpand", "n",false,models.VertualIPS_Noexpand_Desc )
	pf.StringVarP(&VirtualIPS.Vhid, "vhid", "v","",models.VertualIPS_Vhid_Desc )
	pf.StringVar(&VirtualIPS.Advbase, "advbase","",models.VertualIPS_Advbase_Desc )
	pf.StringVar(&VirtualIPS.Advskew, "advskew","",models.VertualIPS_Advskew_Desc )
	pf.StringVarP(&VirtualIPS.Password, "password", "p","",models.VertualIPS_Password_Desc )
}
func createVirtualIPS()  {
	jsonReq, _ := json.Marshal(VirtualIPS)
	res := functions.JsonOutput(jsonReq)
	fmt.Println(res)
	req, err := http.NewRequest("POST", functions.ViperReadConfig().Host+"v1/firewall/virtual_ip", bytes.NewBuffer(jsonReq))
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
	fmt.Println("response body : ", resp.Body)
	fmt.Println("response Headers : ", resp.Header)
}
//pfcli virtualIps create --mode ipalias --interface 172.16.77.239 --subnet 172.16.77.239/32 --password testpass --descr This_is_a_virtual_IP_added_via_pfSense_API
