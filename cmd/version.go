package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "this command will return you CLI version",

		Run: func(cmd *cobra.Command, args []string) {
			version := "1.0.4"
			fmt.Println("pfcli version: ", version)
		},
	}
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

