package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func errchk(i error) {
	if i != nil {
		log.Fatal(i)
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	errchk(err)
	scanner := bufio.NewScanner(file)
	newfilename := strings.TrimSuffix(os.Args[1], ".soh") + ".go"
	newfile, err := os.Create(newfilename)
	_ = newfile
	errchk(err)
	newfile.WriteString("package main\n")

	for scanner.Scan() {
		parts1 := strings.SplitN(scanner.Text(), " ", 2)
		if parts1[0] == "P" {
		}
	}

	for scanner.Scan() {
		parts1 := strings.SplitN(scanner.Text(), " ", 2)
		fmt.Println(parts1[0])
	}
}
