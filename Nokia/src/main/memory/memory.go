package main

import (
	"fmt"
	"runtime"
	_ "net/http/pprof"
	"net/http"
)

func main() {
	fmt.Println(runtime.NumCPU())
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
	
	fmt.Println("count of goroutines ", runtime.NumGoroutine())
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
