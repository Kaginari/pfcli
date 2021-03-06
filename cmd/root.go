package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pfcli",
	Short: `welcome to pfsense client `,
	Long: "pfcli is a go client (CLI) to  pfsense ",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child cmd to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func AddCommands() {


	// Root Commands

	rootCmd.AddCommand(versionCmd())

	rootCmd.AddCommand(NatOneToOneCmd)
	NatOneToOneAddCommands()

	rootCmd.AddCommand(FirewallRuleCmd)
	FirewallRuleAddCommands()

	rootCmd.AddCommand(VirtualIpsCmd)
	VirtualIpsAddCommands()

	rootCmd.AddCommand(InterfaceVLANCmd)
	InterfaceVLANAddCommands()

	rootCmd.AddCommand(interfaceCmd)
	InterfaceAddCommande()

	rootCmd.AddCommand(ConfigCmd)
	ConfigAddCommande()

}
 // init cli configuration

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pfcli/config.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".pfcli" (without extension).
		viper.AddConfigPath("$HOME/.pfcli/")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}


