package myUrl

import (
	"fmt"
	"io/ioutil"
	"myGo/config"
	"net/http"
	"net/url"
	"os"
	"time"
)

func Login(cfg config.Config) error {
	//data := "username%3Dwperrigo%2C%0Apassword%3D" + cfg.Passwd
	/*data := url.Values{}
	data.Add("username", cfg.Username)
	data.Add("password", cfg.Passwd)
	var pwArr [1]string
	pwArr[0] = config.Config.String(cfg.Passwd)*/
	Client := &http.Client{}
	requestURL := fmt.Sprintf(cfg.SplunkBaseUrl + cfg.SplunkLoginUrl)
	params := url.Values{}
	params.Add("username", cfg.Username)
	params.Add("password", cfg.Passwd)
	params.Encode()
	val := url.Values(params)
	req, err := http.PostForm(requestURL, val) // URL-encoded payload
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	//req.Header.Add("Access-Control-Allow-Credentials", "true")
	resp, err := Client.Do(req.Request)
	if err != nil {
		fmt.Printf("client: could not create response: %s\n", err)
		os.Exit(2)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("response", err)
	}
	fmt.Print("print", b)

	/*resp map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&resp)
	*/
	fmt.Print(resp.Body)

	return nil
}

func Call(cfg config.Config, method string) error {
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest(method, cfg.SplunkBaseUrl+cfg.SplunkLoginUrl, nil)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	req.SetBasicAuth(cfg.Username, cfg.Passwd)
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	fmt.Print(response)
	defer response.Body.Close()
	return nil
}
