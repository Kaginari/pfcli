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
	"encoding/json"
	"fmt"
	"github.com/Kaginari/pfcli/lib"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		config := lib.GetConfig()
		pfClient:=config.Context()
		service:=lib.BackupServiceConstruct(pfClient)
		res , err := service.Save()
		if err != nil {
			fmt.Println("un error est occurred")
			// TODO FIN BETTER WAY TO HANDLE ERRORS
		}
		jsonRes, _ := json.Marshal(res.Date)
		ioutil.WriteFile("big_marhsall.json", jsonRes, os.ModePerm)

		fmt.Println("done")
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}
