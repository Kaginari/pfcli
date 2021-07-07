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
	"github.com/Kaginari/pfcli/lib"
	"github.com/Kaginari/pfcli/models"
	"github.com/spf13/cobra"
)

// addCmd represents the add command

var SetCmd = &cobra.Command{
	Use:   "set",
	Run: func(cmd *cobra.Command, args []string) {
		list :=[]string{configmodel.Host,configmodel.ClientId,configmodel.ClientToken}
		lib.ViperAddConfig(list)

		fmt.Println(lib.ViperReadConfig())
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