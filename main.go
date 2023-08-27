package main

import (
	"EasierConnect/core"
	"flag"
	"fmt"
	"log"
	"runtime"
)

func main() {
	// CLI args
	host, port, username, password, socksBind,debugDump := "", 0, "", "", "", false
	flag.StringVar(&host, "server", "oa1.anhuitelecom.com", "EasyConnect server address (e.g. oa1.anhuitelecom.com, oa2.anhuitelecom.com)")
	flag.StringVar(&username, "username", "", "Your username")
	flag.StringVar(&password, "password", "", "Your password")
	flag.StringVar(&socksBind, "socks-bind", ":1088", "The addr socks5 server listens on (e.g. 127.0.0.1:1088)")
	flag.IntVar(&port, "port", 443, "EasyConnect port address (e.g. 443)")
	flag.BoolVar(&debugDump, "debug-dump", false, "Enable traffic debug dump (only for debug usage)")
	flag.Parse()

	if host == "" || username == "" || password == "" {
		log.Fatal("Missing required cli args, refer to `EasierConnect --help`.")
	}
	server := fmt.Sprintf("%s:%d", host, port)

	client := core.NewEasyConnectClient(server)

	var ip []byte
	var err error

	ip, err = client.Login(username, password)
	if err == core.ERR_NEXT_AUTH_SMS {
		fmt.Print(">>>Please enter your sms code<<<:")
		smsCode := ""
		fmt.Scan(&smsCode)

		ip, err = client.AuthSMSCode(smsCode)
	}

	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Login success, your IP: %d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])

	client.ServeSocks5(socksBind, debugDump)

	runtime.KeepAlive(client)
}
