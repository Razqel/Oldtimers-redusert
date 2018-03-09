package main

import (
	"os"
	"log"
	"fmt"
	"bufio"
	"strings"
	"os/signal"
	"syscall"
	"time"
)
/*
*
*
KJØRES SLIK:
1.
go run addtofile.go numbers.txt 1337 1337 (VELG 2 TALL)

2.
go run sumtofile.go numbers.txt (PLUSSER TALL 1 OG TALL 2 FRA ISTAD)

3.
go run addtofile.go numbers.txt (PRINTER UT TOTALSUMMEN FRA PLUSSINGA)


--------------------------------------------------------------------------------------------------------
Til oppgave 3: Bruk Channels i oppgave a.
I Oppgave 3b må dere implementere en låsemekanisme slik at kun en prosess om gangen kan
skrive til den delte filen. En mulighet er en .lock file (https://fileinfo.com/extension/lock
(Lenker til en ekstern side.)Lenker til en ekstern side.) - når denne filen eksisterer er ressursen låst.



DETTE ER --PROGRAM A--
HVA GJØR PROGGRAM A? (OG PROGRAM B?)

"POSTER" TALLENE FRA ARGS[2] OG ARGS[3] TIL ANGITT FILAVN I ARGS[1]
"LESER" OGSÅ FIL FOR Å FINNE RESULTAT SOM ER SKREVET AV PROGRAM B
A:1
A:2
B:3
OG SÅ SKAL PROGRAM A KUNNE LESE RESULATET FRA PROGRAM B
*/

func writeNumbers() {
	input := os.Args[1]
	file, err := os.OpenFile(input, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		log.Fatal("Error opening file:", err)
		return
	}
	defer file.Close()

	// Write bytes to file
	input2 := os.Args[2]
	input3 := os.Args[3]

	stringWritten, err := file.WriteString("A:"+ input2 + "\n" + "A:" + input3 + "\n")
	if err != nil {
		log.Fatal("Error writing:", err)
		return
	}
	log.Printf("Wrote %d bytes.\n", stringWritten)
}
func readTotal() {

	input := os.Args[1]
	file, err := os.OpenFile(input, os.O_RDWR|os.O_EXCL, 0666)
	if err != nil {
		log.Fatal("Error opening file:", err)
		return
	}
	defer file.Close()
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines {
		fmt.Println(line)
		if strings.HasPrefix(string(line), "B:") {
			fmt.Println("Program B - Total: " + line[2:])
		}
	}
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

	if len(os.Args) == 2 {
		fmt.Print("Reading file MODE\n")
		readTotal()
	} else if len(os.Args) == 4 {

		fmt.Print("Writing to file MODE\n")
		writeNumbers()
	} else {
		fmt.Println("Need correct amount of arguments. 1 or 3 args.")
	}


}

