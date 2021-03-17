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
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/nt-factory/2021/admin/pfcli/models"
)


var configCmd = &cobra.Command{
	Use:   "config-list",
	Short: "return the config from the file $ HOME / .pfcli.yaml.",

	Run: func(cmd *cobra.Command, args []string) {
		ViperReadConfig()

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func ViperReadConfig() models.Config  {
	var config models.Config
	viper.SetConfigName("pfcli")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	url_pfsense := fmt.Sprintf("%v", viper.Get("url-pfsense"))
	config.UrlPfsense= url_pfsense
	client_id := fmt.Sprintf("%v", viper.Get("client-id"))
	config.ClientId= client_id
	client_token := fmt.Sprintf("%v", viper.Get("client-token"))
	config.ClientToken=client_token
	fmt.Println(config)
	return config
}
