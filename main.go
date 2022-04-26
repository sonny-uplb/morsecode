/*
Created by RSM
This program converts Morse code text into English
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	morse = make(map[string]string)
)

// International Morse Code map
func init() {
	morse[".-"] = "a"
	morse["-..."] = "b"
	morse["-.-."] = "c"
	morse["-.."] = "d"
	morse["."] = "e"
	morse["..-."] = "f"
	morse["--."] = "g"
	morse["...."] = "h"
	morse[".."] = "i"
	morse[".---"] = "j"
	morse["-.-"] = "k"
	morse[".-.."] = "l"
	morse["--"] = "m"
	morse["-."] = "n"
	morse["---"] = "o"
	morse[".--."] = "p"
	morse["--.-"] = "q"
	morse[".-."] = "r"
	morse["..."] = "s"
	morse["-"] = "t"
	morse["..-"] = "u"
	morse["...-"] = "v"
	morse[".--"] = "w"
	morse["-..-"] = "x"
	morse["-.--"] = "y"
	morse["--.."] = "z"
	morse[".----"] = "1"
	morse["..---"] = "2"
	morse["...--"] = "3"
	morse["....-"] = "4"
	morse["....."] = "5"
	morse["-....."] = "6"
	morse["--..."] = "7"
	morse["---.."] = "8"
	morse["----."] = "9"
	morse["-----"] = "0"
	morse[".-.-.-"] = "."
	morse["--..--"] = ","
	morse[".----."] = "'"
	morse["..--.."] = "?"
	morse["-.-.--"] = "!"
	morse["/"] = " "
}

func main() {

	// check for arguments
	if len(os.Args[1:]) < 4 {
		fmt.Println("Usage: program -f inputfile -o outfile")
		os.Exit(1)
	}

	var inputfile, outputfile string
	flag.StringVar(&inputfile, "f", "test.txt", "Specify input file to read")
	flag.StringVar(&outputfile, "o", "output.txt", "Specify output file to write")

	flag.Parse()

	// get file to read
	readfile, err := os.Open(inputfile)
	if err != nil {
		fmt.Println("Error opening ", inputfile, "for reading")
	}
	defer readfile.Close()

	// create file to write
	writefile, err := os.Create(outputfile)
	if err != nil {
		fmt.Println("Error opening ", outputfile, "for writing")
	}
	defer writefile.Close()

	scanner := bufio.NewScanner(readfile)
	scanner.Split(bufio.ScanLines) // scanned per line to mark new line when writing to file

	for scanner.Scan() {
		// split into words
		res := strings.Split(scanner.Text(), " ")

		// convert morse code to letters and numbers
		for _, value := range res {
			_, err := writefile.WriteString(morse[value])
			if err != nil {
				return
			}
		}

		// add new line
		_, err := writefile.WriteString("\n")
		if err != nil {
			return
		}
	}
	fmt.Println("Execution finished, kindly check output file")
}
