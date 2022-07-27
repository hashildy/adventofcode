package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	// Create slices of type string to store the values for later manipulation
	text := []string{}
	direction := []string{}
	num_text := []string{}

	//  Create struct for storing data of different types together
	//type submarine struct {
	//	position string
	//	depth int
	//}

	// Create slice of type int to convert from num_text
	num := []int{}

	// read the file line by line
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	// For loop which iterates through each item in the Scanner text and breaks down each word in the line
	for scanner.Scan() {
		text = append(text, scanner.Text())
		if err != nil {
			log.Fatal("Scanner be dead", err)
		}
	}
	//Separate the direction and the number values into their own respective string slices
	for i, v := range text {
		if i%2 == 0 {
			direction = append(direction, v)
		} else {
			num_text = append(num_text, v)
		}
	}

	// print the direction & num to the stdout
	//fmt.Println("The direction is: ", direction)
	//fmt.Println("The number of steps is: ", num_text)

	//Convert existing num_text slice into a slice of type int

	for _, n := range num_text {

		tmp, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal("Scanner be dead", err)
		}
		num = append(num, tmp)
	}
	// Test printing the slice of type num
	//fmt.Println(num)

	// Calculate the position & depth by iterating through the for loop and adding up the appropriate values
	pos := 0
	depth := 0

	for a := 0; a < len(direction); a++ {
		if direction[a] == "forward" {
			pos = pos + num[a]
		} else if direction[a] == "up" {
			depth = depth - num[a]
		} else if direction[a] == "down" {
			depth = depth + num[a]
		}
	}
	// Test printing out final result
	fmt.Println("Final position is: ", pos)
	fmt.Println("Final depth is: ", depth)

	// Calculate the final answer
	answer := pos * depth
	fmt.Println("Final answer is position times depth: ", answer)

}
