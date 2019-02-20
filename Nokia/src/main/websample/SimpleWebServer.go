package main

import (
	"net/http"
	"fmt"
//	"strings"
	"main/util"
//	"mime/multipart"
//	"crypto/tls"
//	"crypto/x509"
//	"flag"
//	"io/ioutil"
//	"log"
//	"io"
//	"bytes"
)

//var (
//	certFile = flag.String("cert", "someCertFile", "A PEM eoncoded certificate file.")
//	keyFile  = flag.String("key", "someKeyFile", "A PEM encoded private key file.")
//	caFile   = flag.String("CA", "someCertCAFile", "A PEM eoncoded CA's certificate file.")
//)

func handleRequest(w http.ResponseWriter, res *http.Request){
	w.Header()
	w.Write([]byte ("Something is working atleast"))
	w.WriteHeader(1)

//	testParams := make(map[string]string)
//	
//	ctype := res.Header.Get("Content-Type")
//	result := strings.Split(ctype, ";")
//	for i := range result {
//		if strings.Contains(result[i], "boundary="){
//			testParams["boundary"] = strings.Split(result[i], "boundary=")[1]
//		}
//		if strings.Contains(result[i], "start="){
//			testParams["start"] = strings.Split(result[i], "start=")[1]
//		}
//	}
//	
//	r := strings.NewReader(util.ReadCloserToString(res.Body))
//	reader := NewReader(r, testParams)
//	
//	buf := new(bytes.Buffer)
//	part, err := reader.NextPart()
//	
//	if part == nil || err != nil {
//		fmt.Println("Expected part1")
//		return
//	}
//	buf.Reset()
//	if _, err := io.Copy(buf, part); err != nil {
//		fmt.Println("part 1 copy: %v", err)
//	}

	fmt.Println(res.Header)
	fmt.Println(util.ReadCloserToString(res.Body))
}

func main(){
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
//		InsecureSkipVerify: false,
//	}
//	tlsConfig.BuildNameToCertificate()
//	transport := &http.Transport{TLSClientConfig: tlsConfig}
//	client := &http.Client{Transport: transport}
//	
//	// Do GET something
//	resp, err := client.Get("https://goldportugal.local:8443")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Body.Close()
//
//	// Dump response
//	data, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Println(string(data))
	
	
	
	
//	ln, err := tls.Listen("tcp", *listen, config)
//	
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
//	
//	// Start the HTTPS server
//	http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
//	
//	// Start the HTTP server and redirect all incoming connections to HTTPS
//	http.ListenAndServe(":8080", http.HandlerFunc(redirectToHttps))
}

//func redirectToHttps(w http.ResponseWriter, r *http.Request) {
//    // Redirect the incoming HTTP request. Note that "127.0.0.1:8081" will only work if you are accessing the server from your local machine.
//    fmt.Println("Redirecting to HTTPS")
//    http.Redirect(w, r, "https://127.0.0.1:8080"+r.RequestURI, http.StatusMovedPermanently)
//}
