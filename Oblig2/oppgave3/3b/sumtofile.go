package main

import (
	"os"
	"fmt"
	"log"
	"strings"
	"strconv"
	"bufio"
	"os/signal"
	"syscall"
	"time"
)
/*

KJØRES SLIK:
1.
go run addtofile.go numbers.txt 1337 1337 (VELG 2 TALL)

2.
go run sumtofile.go numbers.txt (PLUSSER TALL 1 OG TALL 2 FRA ISTAD)

3.
go run addtofile.go numbers.txt (PRINTER UT TOTALSUMMEN FRA PLUSSINGA)

 */

func readFile() (total int) {
	input := os.Args[1]
	file, err := os.OpenFile(input, os.O_RDWR | os.O_APPEND | os.O_EXCL, 0666 )

	if err != nil {
		log.Fatal("Open file fail:", err)
		return
	}

	defer file.Close()
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		lines = append(lines, scanner.Text())
	}

	// A: 1
	// A: 2
	// Skal "POSTE" (som Andreas kaller det) "B: 3" til samme fil

	var numbers []int // array med integers (tall, og KUN tall)
	for _, line := range lines {
		if strings.HasPrefix(line, "A:") {
			str, err := strconv.Atoi(line[2:])
			if err != nil {
				log.Panic("cant convert to integer: ", err)
				return
			}
			numbers = append(numbers, str)
		} else {
			log.Fatal("Program A has not provided any numbers")
			return
		}
	}

	for _, sum := range numbers {
		total = total + sum
	}
	fmt.Println("Total:", total)
	return total
}

func writeFile(total int) {
	input := os.Args[1]
	file2, err := os.OpenFile(input, os.O_RDWR | os.O_APPEND | os.O_EXCL, 0666 )
	str := "B:" + strconv.Itoa(total)

	stringWritten, err := file2.WriteString(str + "\n")
	if err != nil {
		log.Fatal("cant write to file:" , err)
		return
	}

	defer file2.Close()
	log.Printf("Wrote %d bytes.\n", stringWritten)
}

func main() {

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		fmt.Printf("Catching SIGNAL: %+v", sig)
		fmt.Println(" 2 second sleep, then exit!")
		time.Sleep(2*time.Second)
		os.Exit(0)
	}()

	var total int

	if len(os.Args) == 2 {
		fmt.Print("Writing file MODE\n")
		total = readFile()
		writeFile(total)
	} else {
		fmt.Println("Need correct amount of arguments. 1 args.")
	}





}


