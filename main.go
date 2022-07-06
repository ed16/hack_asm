package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//Get filename from CLI args
	var filename string

	if len(os.Args) > 0 {
		filename = os.Args[1]
	} else {
		fmt.Println("Please, specify the file *.hack to be assembled")
		return
	}

	i := strings.Index(filename, ".")

	if i == -1 {
		fmt.Println("Please, specify the file *.hack to be assembled")
		return
	}

	if filename[i:] != ".hack" {
		fmt.Println("Please, specify the file *.hack to be assembled")
		return
	}

	fmt.Println(filename)
}
