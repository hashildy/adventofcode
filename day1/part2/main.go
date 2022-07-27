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

	// C++ style for loop to iterate through every 3rd value in our slice and determine if the index is divisible by 3 evenly
	// We also introduce a count here to collect the number of times the next sum is greater than the last sum
	count := 0

	for i := 3; i < len(data); i++ {
		sum1 := 0
		sum2 := 0
		//fmt.Println("This one is divisible by 3!")

		sum1 += data[i-1] + data[i-2] + data[i-3]
		sum2 += data[i] + data[i-1] + data[i-2]

		if sum1 < sum2 {
			fmt.Println("Sum has increased!", sum2)
			count++
		}
		//}
	}

	//Print out the number of times an increase happened by referencing the count variable
	fmt.Println("The number of times an increase happened: ", count)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
