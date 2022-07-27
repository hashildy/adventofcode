package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// open the file positions.txt to read its' contents
	file, err := os.Open("positions.txt")
	if err != nil {
		fmt.Println("Error encountered: ", err)
		os.Exit(1)
	}

	//Don't forget to close the file
	defer file.Close()

	//variable declaration
	var pos, depth, aim, num int
	var direction string

	// read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// read in each line while there are still lines to be read
		text := strings.Fields(scanner.Text())

		//put the first field into the direction variable
		direction = text[0]

		//convert string to int & put value into num variable
		num, err = strconv.Atoi(text[1])

		// perform our program logic

		if direction == "forward" {
			pos += num
			depth += aim * num
		} else if direction == "up" {
			aim -= num
		} else if direction == "down" {
			aim += num
		}
		// testing output
		//fmt.Println("direction: ", direction)
		//fmt.Println("position: ", num)
	}

	// Test printing out final result
	fmt.Println("Final position is: ", pos)
	fmt.Println("Final depth is: ", depth)

	// Calculate the final answer
	answer := pos * depth
	fmt.Println("Final answer is position times depth: ", answer)
}
