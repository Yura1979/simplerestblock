package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Yura1979/simplerestblock/utils"
)

func main() {
	storageSN, err := utils.GetInput("Enter storage serial number: ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Storage serial number is:", storageSN)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	storageIP := "172.21.81.1"
	baseURL := "https://" + storageIP + "/ConfigurationManager/v1/objects/"
	// urlEnd := "/ConfigurationManager/configuration/version"
	// urlEnd := "storages/instance"
	urlEnd := "storage-summaries/instance"
	url := baseURL + urlEnd
	user := "maintenance"
	password := "raid-maintenance"
	fmt.Println(url)

	checkStorage("https://"+storageIP+"/ConfigurationManager/v1/objects/storages", storageIP)

	// resp, err := http.Get(url)
	getStorageInfo(url, user, password)
}

func checkStorage(url, ip string) bool {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	return true
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
	defer resp.Body.Close()
	fmt.Println(string(body))
}
