package services

import (
	"fmt"
	"github.com/spf13/viper"
	"gitlab.com/nt-factory/2021/admin/pfcli/models"
)

func ViperReadConfig() models.Config  {
	var config models.Config
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	host := fmt.Sprintf("%v", viper.Get("host"))
	config.Host= host
	client_id := fmt.Sprintf("%v", viper.Get("client-id"))
	config.ClientId= client_id
	client_token := fmt.Sprintf("%v", viper.Get("client-token"))
	config.ClientToken=client_token
	return config
}
func ViperAddConfig(args []string)  {
	if args[0]==""{args[0]=fmt.Sprintf("%v", viper.Get("host"))}
	if args[1]==""{args[1]=fmt.Sprintf("%v", viper.Get("client-id"))}
	if args[2]==""{args[2]=fmt.Sprintf("%v", viper.Get("client-token"))}
	defaults := map[string]string{"host": args[0], "client-id": args[1] ,"client-token": args[2]}
	for k,v :=range defaults {
		viper.Set(k, v)
	}
	viper.WriteConfig()

}
