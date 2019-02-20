package util

import (
	"bytes"
	xj "github.com/basgys/goxml2json"
	"io"
	"strings"
)

func ReadCloserToString(data io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(data)
	return buf.String()
}

func XmltoJson() string {
	xml := strings.NewReader(`<?xml version="1.0" encoding="UTF-8"?><app version="1.0"><request>1111</request></app>`)

	json, err := xj.Convert(xml)
	if err != nil {
		
	}

	return json.String()
}
