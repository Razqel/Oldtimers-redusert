package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"log"
)


func main() {
	input := os.Args [1]
	b, err := ioutil.ReadFile(input) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	testString := str
	testArray := strings.Fields(testString)

	var sum int
	for _, v := range testArray {
		number1,_ := strconv.Atoi(v)//convert string array to int
		fmt.Println(number1)

		sum = sum + number1 // add numbers from file.

	}
	fmt.Print("Result of the numbers is: ", sum, "\n")

	intToString := strconv.Itoa(sum)
	// Open a new file for writing only
	file, err := os.OpenFile(
		input,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write bytes to file

	bytesWritten, err := file.WriteString(intToString)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
}



