package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"runtime/debug"
)

func main() {
	http.HandleFunc("/Hello", sayHello)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		panic(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	debug.PrintStack()
	fmt.Println(runtime.NumGoroutine())
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}
