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
package _interface

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.com/nt-factory/2021/admin/pfcli/functions"
	"gitlab.com/nt-factory/2021/admin/pfcli/models"
	"gitlab.com/nt-factory/2021/admin/pfcli/lib"
)

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Run: func(cmd *cobra.Command, args []string) {
		config:=lib.GetConfig()
		pfClient:=config.Context()
		service:=lib.InterfaceServiceConstruct(pfClient)
		res,err:=service.Create(InterfaceModel)
		if err!=nil{
			fmt.Println("un error est occurred")
		}
		jsonRes, _ := json.Marshal(res)
		rest := functions.JsonOutput(jsonRes)

		fmt.Println(rest)
	},
}

func init() {
	createFlag()
}
var InterfaceModel models.Interface
func createFlag()  {
	pf := CreateCmd.PersistentFlags()
	pf.StringVarP(&InterfaceModel.If, "if", "i", "", models.InerfaceIfDescr)
	pf.BoolVar(&InterfaceModel.Enable, "enable",false, models.InerfaceEnableDescr)
	pf.StringVarP(&InterfaceModel.Descr, "descr", "d", "", models.Inerface_descr_Descr)
	pf.StringVarP(&InterfaceModel.Subnet, "subnet", "s", "", models.InerfaceSubnetDescr)
	pf.StringVarP(&InterfaceModel.Type, "type", "t", "", models.InerfaceTypeDescr)
	pf.BoolVar(&InterfaceModel.Blockbogons, "blockbogons",  true, models.InerfaceBlockbogonsDescr)
}
