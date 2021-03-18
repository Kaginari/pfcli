package functions


import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gitlab.com/nt-factory/2021/admin/pfcli/models"
)

func CheckRequiredFlags(flags *pflag.FlagSet) error {
	requiredError := false
	flagName := ""

	flags.VisitAll(func(flag *pflag.Flag) {
		requiredAnnotation := flag.Annotations[cobra.BashCompOneRequiredFlag]
		if len(requiredAnnotation) == 0 {
			return
		}

		flagRequired := requiredAnnotation[0] == "true"

		if flagRequired && !flag.Changed {
			requiredError = true
			flagName = flag.Name
		}
	})

	if requiredError {
		return errors.New("Required flag `" + flagName + "` has not been set")
	}

	return nil

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
	return config
}