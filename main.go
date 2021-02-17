package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RepoReleases struct {
	Array []RepoRelease
}

type RepoRelease struct {
	url    string             `json:"url"`
	assets []RepoReleaseAsset `json:"assets"`
}

type RepoReleaseAsset struct {
	browserDownloadUrl string `json:"browser_download_url"`
}

func check(e error, msg string) {
	if e != nil {
		if msg != "" {
			fmt.Println(msg)
		} else {
			panic(e)
		}
	}
}

func prepareMkcert() bool {
	res, _ := http.Get("https://api.github.com/repos/FiloSottile/mkcert/releases")
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var arbitrary_json map[string]interface{}
	json.Unmarshal([]byte(body), &arbitrary_json)

	for key, val := range arbitrary_json {
		fmt.Print(key, val)
	}

	return true
}

func main() {

	prepareMkcert()

	//	var hostName string
	//	var rootPath string
	//
	//	fmt.Print("Enter the domain name: ")
	//	fmt.Scan(&hostName)
	//
	//	fmt.Print("Enter project root folder path: ")
	//	fmt.Scan(&rootPath)
	//
	//
	//	newVHostTag := fmt.Sprintf(`
	//<VirtualHost *:80>
	//	DocumentRoot "%s"
	//	ServerName %s
	//</VirtualHost>
	//
	//<VirtualHost *:443>
	//	DocumentRoot "%s"
	//	ServerName %s
	//	SSLEngine on
	//	SSLCertificateFile "F:\Devtools\certs\nova.local.pem"
	//	SSLCertificateKeyFile "F:\Devtools\certs\nova.local-key.pem"
	//</VirtualHost>
	//`, rootPath, hostName, rootPath, hostName)
	//
	//	fmt.Print(newVHostTag)

	//hostsFilePath := "./secure-vhost/test.txt"
	//
	//f, _ := os.OpenFile(hostsFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//f.WriteString(newVHostTag)
	//defer f.Close()
}
