package main

import (
	"bufio"
	"fmt"
	"myGo/config"
	"myGo/myUrl"
	"os"
)

var Counter = 0
var reader = bufio.NewReader(os.Stdin)

func main() {
	cfg, err := config.CreateConfig()
	if err != nil {
		fmt.Println(err)
	}
	if cfg.Username == "" {
		fmt.Println("Time to authenticate:")
		fmt.Println("---------------------")
		fmt.Println("Enter LDAP Username: ")
		reader := bufio.NewReader(os.Stdin)
		cfg.Username, err = reader.ReadString('\n')
	}
	if cfg.Passwd == "" {
		fmt.Println("Enter LDAP Password: ")
		cfg.Passwd, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(cfg.Username)
	}
	myUrl.Login(cfg)
}
