package main

import (
"fmt"
"os"
"strconv"
)

func foo(c chan int, someValue int, someValue2 int){
c <- someValue + someValue2
}

func main (){
input := os.Args[1]
input2 := os.Args[2]
tall1,_ :=strconv.Atoi(input)
tall2,_ :=strconv.Atoi(input2)
fooVal := make(chan int)
go foo(fooVal, tall2, tall1)
v1 := <-fooVal

fmt.Println(v1)
}