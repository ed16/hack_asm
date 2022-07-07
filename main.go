package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//Get filename from CLI args
	var filename string

	if len(os.Args) > 0 {
		filename = os.Args[1]
	} else {
		fmt.Println("Please, specify the file *.asm to be assembled")
		return
	}

	i := strings.Index(filename, ".")

	if i == -1 {
		fmt.Println("Please, specify the file *.asm to be assembled")
		return
	}

	if filename[i:] != ".asm" {
		fmt.Println("Please, specify the file *.asm to be assembled")
		return
	}

	fmt.Println(filename)

	//Read *.asm file
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	var asm_array []string
	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()
		line = strings.ReplaceAll(line, " ", "")
		if len(line) > 1 && line[:2] != "//" {
			asm_array = append(asm_array, line)
			fmt.Println(line)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i, line := range asm_array {

	}
}
