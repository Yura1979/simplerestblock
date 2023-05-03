package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	storage := "172.21.81.1"
	baseURL := "https://" + storage + "/ConfigurationManager/v1/objects/"
	// urlEnd := "/ConfigurationManager/configuration/version"
	urlEnd := "/storages"
	url := baseURL + storage + urlEnd
	fmt.Println(url)

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
