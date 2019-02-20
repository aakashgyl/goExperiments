package main

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	user          = "Nemuadmin"
	password      = "nemuuser"
	loginURL      = "https://192.168.255.129/protected/login.cgi"
	sshServiceURL = "https://192.168.255.129/protected/sshservice.html"
	baseURL       = "https://192.168.255.129/protected/"
)

//	flag.Parse()
//
//	// Load client cert
//	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Load CA cert
//	caCert, err := ioutil.ReadFile(*caFile)
//	if err != nil {
//		log.Fatal(err)
//	}
//	caCertPool := x509.NewCertPool()
//	caCertPool.AppendCertsFromPEM(caCert)
//
//	// Setup HTTPS client
//	tlsConfig := &tls.Config{
//		Certificates: []tls.Certificate{cert},
//		RootCAs:      caCertPool,
//	}
//	tlsConfig.BuildNameToCertificate()
//	transport := &http.Transport{TLSClientConfig: tlsConfig}

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{Transport: tr}
	cookie := login(c)
	fmt.Println(cookie)
	tokens := sshService(c, cookie)
	enableDisableSSHLogin(c, cookie, tokens, "enable")
}

func login(c *http.Client) string {
	userpass := user + ":" + password
	encodedCreds := base64.URLEncoding.EncodeToString([]byte(userpass))
	fmt.Println(encodedCreds)
	u, err := url.Parse(loginURL)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("Username", user)
	q.Set("Password", password)
	q.Set("EncodedCreds", encodedCreds)
	q.Set("Validate1", "Validate2")
	u.RawQuery = q.Encode()
	fmt.Println(u)
	req, _ := http.NewRequest("POST", loginURL, strings.NewReader(q.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.Do(req)
	return resp.Header.Get("Set-Cookie")
}

func sshService(client *http.Client, cookie string) map[string]string {
	sshServiceReq, _ := http.NewRequest("GET", sshServiceURL, nil)
	sshServiceReq.Header.Set("Cookie", cookie)
	sshServiceResp, _ := client.Do(sshServiceReq)
	body, _ := ioutil.ReadAll(sshServiceResp.Body)
	tokens := parseSSHResponse(body)
	return tokens
}

func enableDisableSSHLogin(client *http.Client, cookie string, tokens map[string]string, task string) {
	u, err := url.Parse(baseURL + task + "Ssh.cgi")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("stamp", tokens["stamp"])
	q.Set("token", tokens["token"])
	q.Set("frame", "sshservice")
	enableReq, _ := http.NewRequest("POST", baseURL+task+"Ssh.cgi", strings.NewReader(q.Encode()))
	enableReq.Header.Set("Cookie", cookie)
	enableResp, _ := client.Do(enableReq)
	body, _ := ioutil.ReadAll(enableResp.Body)
	fmt.Println(string(body))
}

func parseSSHResponse(data []byte) map[string]string {
	valuemap := make(map[string]string)
	datareader := bytes.NewReader(data)
	htmldata := html.NewTokenizer(datareader)

	for {
		token := htmldata.Next()

		switch {
		case token == html.ErrorToken:
			// End of the document, we're done
			return valuemap
		case token == html.StartTagToken:
			element := htmldata.Token()

			//	<input type=hidden name=stamp value="1072915964">
			//<input type=hidden name=token value="0e6a988b4284ebd1c883adac48e10234e944074d75dfdf56c4cb03ffe87cb537">

			attribmap := make(map[string]string)
			if element.Data == "input" {
				for _, a := range element.Attr {
					if a.Key == "name" || a.Key == "value" {
						attribmap[a.Key] = a.Val
					}
				}
			}

			if nameval, ok := attribmap["name"]; ok {
				if nameval == "stamp" || nameval == "token" {
					valuemap[nameval] = attribmap["value"]
				}
			}

			if len(valuemap) == 2 {
				return valuemap
			}
		}
	}
}
