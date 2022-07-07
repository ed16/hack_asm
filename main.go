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

	for line_num, line := range asm_array {
		// Assign address to var if it appears for the first time
		if line[0:1] == "@" {
			if _, ok := var_list[line[1:]]; !ok {
				var_list[line[1:]] = freeRAM
				freeRAM++
			}
		}
		bin_array = addInstruction(bin_array, line, var_list, line_num)
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

	var_list["SP"] = 0
	var_list["LCL"] = 1
	var_list["ARG"] = 2
	var_list["THIS"] = 3
	var_list["THAT"] = 4
	var_list["SCREEN"] = 16384
	var_list["KBD"] = 24576

	var_list["R0"] = 0
	var_list["R1"] = 1
	var_list["R2"] = 2
	var_list["R3"] = 3
	var_list["R4"] = 4
	var_list["R5"] = 5
	var_list["R6"] = 6
	var_list["R7"] = 7
	var_list["R8"] = 8
	var_list["R9"] = 9
	var_list["R10"] = 10
	var_list["R11"] = 11
	var_list["R12"] = 12
	var_list["R13"] = 13
	var_list["R14"] = 14
	var_list["R15"] = 15

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
			continue
		}

		asm_array = append(asm_array, line)
		n++
		//fmt.Println(line)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return asm_array
}

func addInstruction(bin_array []string, line string, var_list map[string]int, line_num int) []string {

	var instruction string
	var addrDec int
	var dest string
	var comp string
	var jump string

	if line[0:1] == "@" {
		// A_COMMAND
		if val, err := strconv.Atoi(line[1:]); err == nil {
			addrDec = val
		} else {
			addrDec = int(var_list[line[1:]])
		}

		addrBin := fmt.Sprintf("%b", addrDec)
		instruction = fmt.Sprintf("%016s", addrBin)
	} else {
		// C_COMMAND dest=comp;jump

		i := strings.Index(line, "=")
		if i != -1 {
			dest = line[:i]
		} else {
			i = 0
		}

		j := strings.Index(line, ";")
		if j != -1 {
			comp = line[i:j]
			jump = line[j+1:]
		} else {
			comp = line[i+1:]
		}

		instruction = "111" + getCompCode(comp, line_num) + getDestCode(dest, line_num) + getJumpCode(jump, line_num)

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

func getCompCode(comp string, line_num int) (compBin string) {

	switch comp {
	case "0":
		compBin = "0101010"
	case "1":
		compBin = "0111111"
	case "-1":
		compBin = "0111010"
	case "D":
		compBin = "0001100"
	case "A":
		compBin = "0110000"
	case "M":
		compBin = "1110000"
	case "!D":
		compBin = "0001101"
	case "!A":
		compBin = "0110001"
	case "!M":
		compBin = "1110001"
	case "-D":
		compBin = "0001111"
	case "-A":
		compBin = "0110011"
	case "-M":
		compBin = "1110011"
	case "D+1":
		compBin = "0011111"
	case "A+1":
		compBin = "0110111"
	case "M+1":
		compBin = "1110111"
	case "D-1":
		compBin = "0001110"
	case "A-1":
		compBin = "0110010"
	case "M-1":
		compBin = "1110010"
	case "D+A":
		compBin = "0000010"
	case "D+M":
		compBin = "1000010"
	case "D-A":
		compBin = "0010011"
	case "D-M":
		compBin = "1010011"
	case "A-D":
		compBin = "0000111"
	case "M-D":
		compBin = "1000111"
	case "D&A":
		compBin = "0000000"
	case "D&M":
		compBin = "1000000"
	case "D|A":
		compBin = "0010101"
	case "D|M":
		compBin = "1010101"
	default:
		log.Fatal("Syntax error: Wrong comp command \"", comp, "\" Line: ", line_num)
	}

	return compBin

}

func getDestCode(dest string, line_num int) (destBin string) {

	switch dest {
	case "":
		destBin = "000"
	case "M":
		destBin = "001"
	case "D":
		destBin = "010"
	case "MD":
		destBin = "011"
	case "A":
		destBin = "100"
	case "AM":
		destBin = "101"
	case "AD":
		destBin = "110"
	case "AMD":
		destBin = "111"
	default:
		log.Fatal("Syntax error: Wrong dest command \"", dest, "\" Line: ", line_num)
	}

	return destBin

}

func getJumpCode(jump string, line_num int) (jumpBin string) {

	switch jump {
	case "":
		jumpBin = "000"
	case "JGT":
		jumpBin = "001"
	case "JEQ":
		jumpBin = "010"
	case "JGE":
		jumpBin = "011"
	case "JLT":
		jumpBin = "100"
	case "JNE":
		jumpBin = "101"
	case "JLE":
		jumpBin = "110"
	case "JMP":
		jumpBin = "111"
	default:
		log.Fatal("Syntax error: Wrong jump command \"", jump, "\" Line: ", line_num)
	}

	return jumpBin
}
