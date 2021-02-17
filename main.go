package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"sync"
)

type WriteCounter struct {
	Total uint64
}

type RepoReleases struct {
	Collection []RepoRelease
}

type RepoRelease struct {
	Url    string             `json:"url"`
	Assets []RepoReleaseAsset `json:"assets"`
}

type RepoReleaseAsset struct {
	BrowserDownloadUrl string `json:"browser_download_url"`
}

type XamppSecureVhost struct {
	domainName, projectRoot string
}

func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 50))

	fmt.Printf("\rDownloading mkcert binary... %s", humanize.Bytes(wc.Total))
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func DownloadFile(url string, filepath string) error {
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	fmt.Println()

	err = os.Rename(filepath+".tmp", filepath)
	if err != nil {
		return err
	}

	return nil
}

func binaryExists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

var sudoWarningOnce sync.Once

func commandWithSudo(cmd ...string) *exec.Cmd {
	if u, err := user.Current(); err == nil && u.Uid == "0" {
		return exec.Command(cmd[0], cmd[1:]...)
	}
	if !binaryExists("sudo") {
		sudoWarningOnce.Do(func() {
			log.Println(`Warning: "sudo" is not available, and mkcert is not running as root. The (un)install operation might fail. ⚠️`)
		})
		return exec.Command(cmd[0], cmd[1:]...)
	}
	return exec.Command("sudo", append([]string{"--prompt=Sudo password:", "--"}, cmd...)...)
}

func PrepareMkcert() bool {
	//osType := runtime.GOOS

	stores := os.Getenv("TRUST_STORES")

	fmt.Println(stores)

	_, err := exec.LookPath("./mkcert")
	panic(err)
	//if err != nil {
	//	res, _ := http.Get("https://api.github.com/repos/FiloSottile/mkcert/releases")
	//	defer res.Body.Close()
	//
	//	body, _ := ioutil.ReadAll(res.Body)
	//	releases := make([]RepoRelease, 0)
	//
	//	json.Unmarshal([]byte(body), &releases)
	//	if len(releases) > 0 {
	//		latestRelease := releases[0]
	//		for _, asset := range latestRelease.Assets{
	//
	//			if strings.Contains(asset.BrowserDownloadUrl, osType) {
	//				DownloadFile(asset.BrowserDownloadUrl, "./mkcert")
	//				break
	//			}
	//		}
	//	}
	//}

	return true
}

func main() {

	PrepareMkcert()

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
