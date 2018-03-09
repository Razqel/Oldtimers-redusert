package main

import (
	"os"
	"log"
"../addfile"
)

//Til oppgave 3: Bruk Channels i oppgave a.

//I Oppgave 3b må dere implementere en låsemekanisme slik at kun en prosess om gangen kan
// skrive til den delte filen. En mulighet er en .lock file (https://fileinfo.com/extension/lock
// (Lenker til en ekstern side.)Lenker til en ekstern side.) - når denne filen eksisterer er ressursen låst.
func addToFile() {
	// Open a new file for writing only
	input := os.Args[1]
	file, err := os.OpenFile(input,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write bytes to file
	input2 := os.Args[2]
	input3 := os.Args[3]

	stringWritten, err := file.WriteString(input2 + "\n" + input3)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", stringWritten)
}


func main() {
addToFile()
addfile.sumFromFile()
}


/*
func main() {
	input := os.Args[1]
	file, err := os.OpenFile(input,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlices := []byte(os.Args[2] + os.Args[3])
	bytesWritten, err := file.Write(byteSlices)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Numbers Written %d", bytesWritten)

	content, err := ioutil.ReadFile("t.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File contents %s", content)

}
*/