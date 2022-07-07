package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	var bin_array []string

	//Get filepath from CLI args
	filepath := getFilePath()
	fmt.Println(filepath)

	// Fill predefined variables
	var_list := fillPredifinedVars()
	freeRAM := 16
	// Read file, delete whitespaces and comments, fill labels to the list
	asm_array := readAsmFile(filepath, var_list)

	for _, line := range asm_array {
		// Assign address to var if it appears for the first time
		if line[0:1] == "@" {
			if _, ok := var_list[line[1:]]; !ok {
				var_list[line[1:]] = freeRAM
				freeRAM++
			}
		}
		bin_array = addInstruction(bin_array, line, var_list)
	}

	writeFile(bin_array, filepath)
}

func getFilePath() (filepath string) {
	if len(os.Args) > 0 {
		filepath = os.Args[1]
	} else {
		//fmt.Println("Please, specify the file *.asm to be assembled")
		log.Fatal("Please, specify the file *.asm to be assembled")
	}

	i := strings.Index(filepath, ".")

	if i == -1 {
		log.Fatal("Please, specify the file *.asm to be assembled")
	}

	if filepath[i:] != ".asm" {
		log.Fatal("Please, specify the file *.asm to be assembled")
	}

	return filepath
}

func fillPredifinedVars() map[string]int {
	var_list := make(map[string]int)

	var_list["R1"] = 1
	var_list["R2"] = 2
	return var_list

}

func readAsmFile(filepath string, var_list map[string]int) (asm_array []string) {
	//Read *.asm file
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var n int
		// processing a line
		line := strings.ReplaceAll(scanner.Text(), " ", "")
		if len(line) == 1 {
			log.Fatal("Syntax error in line ", n)
		}
		//Delete full line comments
		if len(line) > 1 && line[:2] == "//" {
			continue
		}
		// Skip empty lines
		if len(line) < 1 {
			continue
		}
		//Delete inline comments
		i := strings.Index(line, "//")
		if i != -1 {
			line = line[:i]
		}

		//Add (labels) to the var_list
		if line[0:1] == "(" {
			i = strings.Index(line, ")")
			line = line[1:i]
			var_list[line] = n + 1
		}

		asm_array = append(asm_array, line)
		n++
		fmt.Println(line)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return asm_array
}

func addInstruction(bin_array []string, line string, var_list map[string]int) []string {

	var instruction string
	var addrDec int

	if line[0:1] == "@" {
		if val, err := strconv.Atoi(line[1:]); err == nil {
			addrDec = val
		} else {
			addrDec = int(var_list[line[1:]])
		}

		addrBin := fmt.Sprintf("%b", addrDec)
		instruction = fmt.Sprintf("%016s", addrBin)
	} else {
		instruction = "1111111111111110"

	}
	fmt.Println(instruction)

	return append(bin_array, instruction)

}

func writeFile(bin_array []string, filepath string) (err error) {
	var (
		file *os.File
	)

	filepath = filepath[:len(filepath)-3] + "txt"

	if file, err = os.Create(filepath); err != nil {
		return
	}

	defer file.Close()

	for _, item := range bin_array {

		_, err := file.WriteString(strings.TrimSpace(item) + "\n")

		if err != nil {
			fmt.Println(err)
			break
		}
	}

	return
}
