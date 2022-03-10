package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	var filename string
	for {
		fmt.Scanln(&filename)
		println("Searching for :", filename)
		println("In file : ", filename)
		file, err := os.Open(filename)

		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var txtlines []string

		for scanner.Scan() {
			txtlines = append(txtlines, scanner.Text())
		}

		file.Close()

		for _, eachline := range txtlines {
			println(eachline)
		}
	}
}
