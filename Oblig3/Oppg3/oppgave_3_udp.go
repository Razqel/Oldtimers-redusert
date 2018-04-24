
//For å kjøre programmen "sudo go run filnavn.go"
//For å teste åpne ny terminal og skriv nc -u 127.0.0.1 17 <- unix/mac

package main

import (
	"fmt"
	"net"
	"os"
)

// Håndterer feil
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {
	/* Setter opp serveren til å lytte på 17*/
	ServerAddr, err := net.ResolveUDPAddr("udp", "localhost:17")
	CheckError(err)
	fmt.Println("listening on :", ServerAddr)

	/* Hører på den valgte porten */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for { //Tar inn melding
		n,addr,err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Received ",string(buf[0:n]), " from ",addr)

		newmessage := ("Quote of the Day")

		// Sender Quoten tilbake til clienten
		ServerConn.WriteToUDP([]byte(newmessage + "\n"),addr)

		if err != nil {
			fmt.Println("Error: ",err)
		}
	}
}