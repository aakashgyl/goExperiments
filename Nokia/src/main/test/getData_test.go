package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetData(t *testing.T){
	data := []byte("[{\"owner\": {\"html_url\": \"200\"}, \"full_name\": \"Go\"}, {\"owner\": {\"html_url\": \"200\"}, \"full_name\": \"Java\"}]")
	datamap := processData(data)
	datamapExpected := make(map[string]string)
	datamapExpected["Go"] = "200"
	datamapExpected["Java"] = "200"
	assert.Equal(t, datamapExpected, datamap)
}