package main

import (
	"fmt"
	"os"
	"strconv"
	"os/signal"
	"syscall"
)

func foo(c chan int, someValue int, someValue2 int) {
	c <- someValue + someValue2
}

func main() {
	input := os.Args[1]
	input2 := os.Args[2]
	tall1, err := strconv.Atoi(input)
	if (err != nil) {
		fmt.Print("Ugyldig input, skriv inn tall ")
	}
	tall2, err := strconv.Atoi(input2)
	if (err != nil) {
		fmt.Print("Ugyldig input, skriv inn tall ")
	}
	fooVal := make(chan int)
	go foo(fooVal, tall2, tall1)
	v1 := <-fooVal
	fmt.Println(v1)

	sigs := make(chan os.Signal, 1)
	done := make(chan bool)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("Venter signal")
	<-done
	fmt.Println("Avslutter")

}
