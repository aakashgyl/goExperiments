package main

import (
	"bytes"
	"compress/gzip"
	"edgene3spm/configs"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"fmt"
)

func main(){
	fmt.Println(GetMimeNotif())
}

func GetMimeNotif() *http.Request {
	req := getRequest("")
	req.RemoteAddr = "127.0.0.1:9091"
	return req
}

func getRequest(format string) *http.Request {
	contentType := "multipart/related"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	err := prepareNotifReq(writer)

	if err != nil {
		return &http.Request{}
	}

	err = loadNotifFile(writer, "feedback.xml")
	if err != nil {
		return &http.Request{}
	}

	req, _ := http.NewRequest("POST", "http://localhost:9001/upload/", body)

	if format == "Video" {
		contentType = "dummy/related"
	}

	req.Header.Add("content-type", strings.Replace(writer.FormDataContentType(), "multipart/form-data", contentType, -1)+"; type=\"text/xml\"; start=\"SOAP-ENV:Envelope\"")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("SOAPAction", "")

	return req

}

func prepareNotifReq(writer *multipart.Writer) error {
	file, err := os.Open("response.xml")

	if err != nil {
		log.Debug("File not found")
		return err
	}
	defer file.Close()

	h := make(textproto.MIMEHeader)
	h.Set("Content-Type", "text/xml; charset=UTF-8")
	h.Set("Content-ID", "<SOAP-ENV:Envelope>")

	part, _ := writer.CreatePart(h)
	io.Copy(part, file)

	return nil
}

func loadNotifFile(writer *multipart.Writer, filepath string) error {
	notifdata, _ := ioutil.ReadFile(filepath)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Type", "text/xml; charset=UTF-8")
	h.Set("Content-ID", "<notificationContent>")

	part, _ := writer.CreatePart(h)
	//		io.Copy(part, file)
	io.Copy(part, strings.NewReader(notifdata))

	if err := writer.Close(); err != nil {
		return err
	}
	return nil
}
