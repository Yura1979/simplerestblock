package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	storage := "172.21.81.1"
	url := "https://" + storage + "/ConfigurationManager/configuration/version"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println(string(body))
}
