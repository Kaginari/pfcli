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
	"gitlab.com/nt-factory/2021/admin/pfcli/models"
	"io"
	"log"
	"os"
	"regexp"
	//"regexp"
	//"strings"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config-list",
	Short: "return the config from the file $ HOME / .pfcli.yaml.",

	Run: func(cmd *cobra.Command, args []string) {
		ReadConfig()

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}


func ReadConfig() models.Config  {
	var config models.Config
	file, err := os.Open("/home/pfcli.yaml")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}

		if n > 0 {
			re := regexp.MustCompile(`\[([^\[\]]*)\]`)
			submatchall := re.FindAllString(string(buf[:n]), -1)
			a:=len(submatchall[0])
			b:=len(submatchall[1])
			c:=len(submatchall[2])
			config.UrlPfsense= string(submatchall[0][1:a-1])
			config.ClientId= string(submatchall[1][1:b-1])
			config.ClientToken=string(submatchall[2][1:c-1])
				fmt.Println(config)
			}
		}

	return config
}