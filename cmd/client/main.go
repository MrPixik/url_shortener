package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
Server testing manually
*/
func main() {
	endpoint := "http://localhost:8080/"

	//Invite in console
	fmt.Println("Enter your URL:")

	//Open standard input from console
	reader := bufio.NewReader(os.Stdin)

	//Read the input string
	longUrl, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	longUrl = strings.TrimSuffix(longUrl, "\n")

	data := []byte(longUrl)

	client := &http.Client{}

	//Creating POST request
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "text/plain")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	//Response processing
	fmt.Println(resp.Status)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

}
