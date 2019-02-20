package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func PrintCWD() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}