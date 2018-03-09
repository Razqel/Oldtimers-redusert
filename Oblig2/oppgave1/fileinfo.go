package main

import (
	"fmt"
	"log"
	"os"
)


func leserFilInfo() {
	input := os.Args[1]
	fi, err := os.Lstat(input)
	if err != nil {
		log.Fatal(err)
	}

	switch mode := fi.Mode(); {

	case mode.IsRegular():
		fmt.Println("Is a regular file")
	case mode.IsDir():
		fmt.Println("Is a directory")
	case mode&os.ModeSymlink != 0:
		fmt.Println("Is symbolic link")
	case mode&os.ModeNamedPipe != 0:
		fmt.Println("Is a named pipe")
	}
}

func isAppendAble() bool {
	input := os.Args[1]
	var appendAble bool // lokal variable som lagrer status på hvorvidt filen er appendable(skrivbar)

	var f,err = os.OpenFile(input, os.O_APPEND|os.O_WRONLY, os.ModePerm) // OpenFile()-funksjonen returnerer to pointers (en til selve filen og en til errorStatus)
	if err != nil { // sjekker om err(or)-variablen nå inneholder en error - isåfall lagrer vi statusen på hvorvidt fila er appendable(skrivbar) i boolean-variablen 'appendable'
		appendAble = false
	} else {
		appendAble = true
	}
		defer f.Close() //lukker fila

		return appendAble // returnerer appendable-variablen (som viser hvorvidt status på fila er appendable/skrivbar)

}

func main() {
	input := os.Args[1] 	//input, aka fila vi skal "Scanne"
		var (
		fileInfo os.FileInfo
		err      error
	)


	fileInfo, err = os.Stat(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Information about file:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size(),"- in KB:", fileInfo.Size()/1024,
		"- in MB:", fileInfo.Size()/1048576, "- in GB:", fileInfo.Size()/1073741824)
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	/*isRegular()*/
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Is Append: ", isAppendAble())
	fmt.Println("Last modified:", fileInfo.ModTime())
	leserFilInfo()
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())

}