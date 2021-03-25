package lib

import (
	"fmt"
	"github.com/spf13/viper"
	"gitlab.com/nt-factory/2021/admin/pfcli/models"
	"os"
)

func ViperReadConfig() models.Config  {
	var config models.Config

	config.Host= os.Getenv("PFCLI_HOST")
	config.ClientId= os.Getenv("PFCLI_CLIENT")
	config.ClientToken= os.Getenv("PFCLI_SECRET")
	if (config.Host!="" && config.ClientId!="" && config.ClientToken!=""){
		fmt.Println("this is from env variable")
		return config
	}else {
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
		fmt.Println("this is from copnfg.yaml")
	return config
	}
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
//pfcli config set --host https://51.77.239.61/api/ --client_id 6170695f75736572 --client_token 231364d73da34275d32c9c79ea69f80c