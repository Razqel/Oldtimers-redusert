package main

import (
	"fmt"
)

func IterateOverASCIIStringLiteral(sl string) {
	for _, r := range sl {
		fmt.Printf("%X %c %b\n", r, r, r)
	}
}
func IterateOver128ASCII(){
	for i := 0; i <= 128; i++ {
		fmt.Printf("%2[1]X %3[1]c %8[1]b\n", i)
	}
}
func ExtendedASCIIText(){
	const signs = "\x80\xF7\xBE"
	const dollar = "\x64\x6F\x6C\x6C\x61\x72"


	for i := 0; i < len(signs); i++ {
		fmt.Printf(" %c", signs[i])
	}
		fmt.Printf(" ")

	for i := 0; i < len(dollar); i++ {
		fmt.Printf("%c", dollar[i])
	}
}


func main() {
		fmt.Println("Extended ASCII print")
		IterateOver128ASCII()
		fmt.Println("\n\n\nIterate print")
		IterateOverASCIIStringLiteral("Hello world")
		fmt.Println("\n\n\nExtendASCIIText Dollar")
		ExtendedASCIIText()
}

