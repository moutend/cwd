package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	wd, _ := os.Getwd()

	fmt.Println(filepath.Base(wd))
}
