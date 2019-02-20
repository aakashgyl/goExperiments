package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"
	"io"
)

func main() {
//	var b bytes.Buffer
//	w := gzip.NewWriter(&b)
//	w.Write([]byte("hello, world\n"))
//	w.Close()
//	_ = ioutil.WriteFile("hello_world.txt.gz", b.Bytes(), 0666)
	
	c, _ := ioutil.ReadFile("hello_world.txt.gz")
	d := bytes.NewReader(c)
	r, _ := gzip.NewReader(d)
	io.Copy(os.Stdout, r)
	r.Close()
}