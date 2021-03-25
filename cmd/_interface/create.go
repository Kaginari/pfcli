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
	"github.com/spf13/cobra"
	"gitlab.com/nt-factory/2021/admin/pfcli/models"
	"gitlab.com/nt-factory/2021/admin/pfcli/lib"
)

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Run: func(cmd *cobra.Command, args []string) {
		lib.CreateInterface(Interface_Model)
	},
}

func init() {
	createFlag()
}
var Interface_Model models.Interface
func createFlag()  {
	pf := CreateCmd.PersistentFlags()
	pf.StringVarP(&Interface_Model.If, "if", "i", "", models.InerfaceIfDescr)
	pf.BoolVar(&Interface_Model.Enable, "enable",false, models.InerfaceEnableDescr)
	pf.StringVarP(&Interface_Model.Descr, "descr", "d", "", models.Inerface_descr_Descr)
	pf.StringVarP(&Interface_Model.Subnet, "subnet", "s", "", models.InerfaceSubnetDescr)
	pf.StringVarP(&Interface_Model.Type, "type", "t", "", models.InerfaceTypeDescr)
	pf.BoolVar(&Interface_Model.Blockbogons, "blockbogons",  true, models.InerfaceBlockbogonsDescr)
}
