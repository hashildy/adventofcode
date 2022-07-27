package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// open the file specified in the go command line to read its contents
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error encountered: ", err)
		os.Exit(1)
	}

	//Don't forget to close the file
	defer file.Close()

	// read the file line by line
	scanner := bufio.NewScanner(file)

	// Create a new slice of type int to store the string values once converted
	data := []int{}

	// For loop which iterates through each item in the Scanner text and appends a value to the data slice
	for scanner.Scan() {

		tmp, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Scanner be dead", err)
		}
		data = append(data, tmp)
	}

	// C++ style for loop to iterate through each value in our slice and determine whether or not the value increased or descreased from the previous value
	// We also introduce a count here to collect the number of times a value increased
	count := 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			//fmt.Println("Increased")
			count++
		}
		//fmt.Println("Decreased")
	}

	//Print out the number of times an increase happened by referencing the count variable
	fmt.Println("The number of times an increase happened: ", count)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
