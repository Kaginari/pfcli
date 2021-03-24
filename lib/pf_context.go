package lib

import (
	"bytes"
	"crypto/tls"
	"net/http"
)
type Config struct {
	Host  string `json:"host"`
	ClientId    string `json:"client_id"`
	ClientToken string `json:"client_token"`
}

type PfClient struct {
	http *http.Client
	secret string
	token string
	host string
}
func GetConfig() Config  {
	return Config{
		Host:        ViperReadConfig().Host,
		ClientId:    ViperReadConfig().ClientId,
		ClientToken: ViperReadConfig().ClientToken,
	}
}
func (c *Config) Context()  PfClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	pfClient := PfClient{
		client,
		c.ClientToken,
		c.ClientId,
		c.Host,
	}
	return pfClient
}

func (pf *PfClient) Post(path string , body []byte ) (*http.Response , error) {

	request , err := http.NewRequest("POST", pf.host + path, bytes.NewBuffer(body))
	if err != nil {
		return nil , err
	}
	request.Header.Add("Authorization", pf.token + " "+ pf.secret)
	response , err := pf.http.Do(request)
	if err != nil {
		return nil , err
	}
	return response , nil
}

func (pf *PfClient) Get(path string , body []byte ) (*http.Response , error) {

	request , err := http.NewRequest("GET", pf.host + path, bytes.NewBuffer(body))
	if err != nil {
		return nil , err
	}
	request.Header.Add("Authorization", pf.token + " "+ pf.secret)
	response , err := pf.http.Do(request)
	if err != nil {
		return nil , err
	}
	return response , nil
}
func (pf *PfClient) Delete(path string , body []byte ) (*http.Response , error) {

	request , err := http.NewRequest("DELETE", pf.host + path, bytes.NewBuffer(body))
	if err != nil {
		return nil , err
	}
	request.Header.Add("Authorization", pf.token + " "+ pf.secret)
	response , err := pf.http.Do(request)
	if err != nil {
		return nil , err
	}
	return response , nil
}

