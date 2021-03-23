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
package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.com/nt-factory/2021/admin/pfcli/models"
	"gitlab.com/nt-factory/2021/admin/pfcli/services"
)

// addCmd represents the add command
var SetCmd = &cobra.Command{
	Use:   "set",
	Run: func(cmd *cobra.Command, args []string) {
		list :=[]string{configmodel.Host,configmodel.ClientId,configmodel.ClientToken}
		services.ViperAddConfig(list)

		fmt.Println(services.ViperReadConfig())
	},
}

func init() {
	createFlag()
}
var configmodel models.Config
func createFlag() {
	pf := SetCmd.PersistentFlags()
	pf.StringVar(&configmodel.Host, "host",  "", models.Hoet_descr)
	pf.StringVar(&configmodel.ClientId, "client_id",  "", models.Client_id_descr)
	pf.StringVar(&configmodel.ClientToken, "client_token",  "", models.Client_token_descr)
}