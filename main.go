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
	// urlEnd := "storages/instance"
	urlEnd := "storage-summaries/instance"
	url := baseURL + urlEnd
	user := "maintenance"
	password := "raid-maintenance"
	fmt.Println(url)

	// resp, err := http.Get(url)
	getStorageInfo(url, user, password)
}

func getStorageInfo(url string, user string, password string) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(user, password)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println(string(body))
}
