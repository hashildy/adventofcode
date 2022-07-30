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
	file, err := os.Open("diagnostics.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// variables declarations

	var lines []string
	var gLine []string
	var eLine []string

	// scan in file and place each line into its own array
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// Calling the get CharSum function to process every character of each line and calculate result
	for x := 0; x < 12; x++ {
		a, b := getCharSum(lines, x)
		fmt.Println(a, b)
		if a > b {
			fmt.Println("Zeros win!")
			gLine = append(gLine, "0")
			eLine = append(eLine, "1")
		}
		if b > a {
			fmt.Println("Ones win!")
			gLine = append(gLine, "1")
			eLine = append(eLine, "0")
		}
		fmt.Println(gLine)
	}

	// Calling the calcPower function for Gamma to convert the binary string into a decimal value
	gamma := calcPower(gLine)
	epsilon := calcPower(eLine)
	fmt.Println("Gamma decimal value is: ", gamma)
	fmt.Println("Epsilon decimal value is: ", epsilon)

	// Calculate power consumption of submarine

	answer := gamma * epsilon
	fmt.Println("Final value is: ", answer)

}

func getCharSum(lines []string, char int) (int, int) {
	zeros, ones := 0, 0
	// Looping through each line one at a time
	for _, c := range lines {
		//fmt.Println(i, " =>", string(c[:1]))
		if string(c[char:char+1]) == "0" {
			zeros++
		} else {
			ones++
		}
	}
	return zeros, ones
}

func calcPower(newline []string) int {
	text := strings.Join(newline, "")
	//fmt.Println(text)
	binary, err := strconv.ParseInt(text, 2, 0)
	if err != nil {
		log.Fatal()
	}
	return int(binary)
}
