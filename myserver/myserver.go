package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server")
	http.HandleFunc("/company", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Company handler")
	})
	http.Handle("/", myhandler{})
	//http.HandleFunc("/getform", getForm)
	//http.HandleFunc("/postform", postForm)
	http.HandleFunc("/test", getpostForm)
	http.Handle("/files/", http.StripPrefix("/files", http.FileServer(http.Dir("."))))
	http.ListenAndServe(":8086", nil)
}

type myhandler struct {
}

func (m myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error

	if err!= nil{
		http.Error(w, "met with an error", 1)		//case when some error happened and I cant do anything
	}else{
		fmt.Fprintln(w, "Hello HTTP")
	}
}

func getForm(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/html")
	fmt.Fprintln(w, `<form method="POST" action="/test"><input name="myname"/><input type="submit"/></form>`)
}

func postForm(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Fprintf(w,"Your name is %s", r.PostFormValue("myname"))
}

func getpostForm(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		getForm(w,r)
	}else if r.Method == "POST"{
		postForm(w,r)
	}
}