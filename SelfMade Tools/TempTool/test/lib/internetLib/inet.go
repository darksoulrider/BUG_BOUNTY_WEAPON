package internetlib

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
)

func Check_NsLookup(name string) {
	domain := name
	_ = domain
	wordlistfile := "wordlist.txt"
	wordlist, err := os.ReadFile(wordlistfile)
	if err != nil {
		fmt.Printf("Error reading wordlist file: %v\n", err)
		os.Exit(1)
	}

	subdomains := string(wordlist)
	subdomainsList := strings.Split(subdomains, "\n")

	for _, subdomain := range subdomainsList {
		subdomain := strings.TrimSpace(subdomain)
		host := fmt.Sprintf("%s.%s", subdomain, domain)

		res, err := net.LookupHost(host)
		if err != nil {
			continue
		}
		_ = res
		fmt.Println("up -> ", subdomain)

	}

}

func GetWebSourceCode(host string) {

	res, err := http.Get(host)
	if err != nil {
		fmt.Println("Error Occured: ", err.Error())
		os.Exit(1)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error Occured: ", err.Error())
		os.Exit(1)
	}

	dd := ioutil.WriteFile("index.html", []byte(body), 0644)
	if dd != nil {
		os.Exit(1)
	}

}
