package main

import (
	"bufio"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
	"strings"
)

var mainMenuMsg = [...]string{
	"1. GET at /",
	"2. POST at /",
	"3. POST at /api/shorten",
	"4. GET at /ping",
	"0. Exit",
}

func readInputString(msg string) string {
	//Invite in console
	fmt.Println(msg)

	//Open standard input from console
	reader := bufio.NewReader(os.Stdin)

	//Read the input string
	longUrl, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	longUrl = strings.TrimSuffix(longUrl, "\n")
	return longUrl
}

func printMainMenuMsg() {
	for _, str := range mainMenuMsg {
		fmt.Println(str)
	}
}

func menu(client *resty.Client) {
	var reqUrl requestUrl
	for {
		printMainMenuMsg()
		key := readInputString("Enter the key: ")

		switch key {
		case "1":
			mainPageGetRequest(client, reqUrl)
		case "2":
			mainPagePostRequest(client, reqUrl)
		case "3":
			shortenApiPostRequest(client, reqUrl)
		case "4":
			pingDBRequest(client)
		case "0":
			return
		default:
			continue
		}
	}
}
