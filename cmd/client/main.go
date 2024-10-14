package main

import (
	"bufio"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
	"strings"
)

func getStr(msg string) string {
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

/*
Server testing manually
*/
func main() {
	endpoint := "http://localhost:8999/"

	longUrl := getStr("Enter your URL:")

	data := []byte(longUrl)

	client := resty.New()

	//Creating POST request
	resp, err := client.R().
		//SetHeader("Content-Type", "text/plain").
		SetBody(data).
		Post(endpoint)
	if err != nil {
		panic(err)
	}

	//Response processing
	fmt.Println("Post response status: " + resp.Status())

	fmt.Println(string(resp.Body()))

}
