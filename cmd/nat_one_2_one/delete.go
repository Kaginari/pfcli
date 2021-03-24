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
package nat_one_2_one

import (
	"github.com/spf13/cobra"
	"gitlab.com/nt-factory/2021/admin/pfcli/models"
	"gitlab.com/nt-factory/2021/admin/pfcli/lib"
)

// deleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Run: func(cmd *cobra.Command, args []string) {
		lib.DeleteNat(DeleteModel)
	},
}

func init() {
	pf := DeleteCmd.PersistentFlags()
	pf.StringVarP(&DeleteModel.Id, "id", "i", "", "Specify the rule tracker ID to delete")
	pf.BoolVar(&DeleteModel.Apply, "apply",true, models.NAT121ApplyDesc)
}
var DeleteModel models.DeleteNatOneToOne
