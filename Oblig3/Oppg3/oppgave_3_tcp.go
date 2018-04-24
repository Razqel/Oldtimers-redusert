

// kjør filen med "sudo go run filnavn.go"
//skriv   echo -n "test out the server" | nc localhost 17   på mac/unix for å teste

package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":17")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g. forbindelse avbrytet
			continue }
		go handleConnection(conn) // en forbindelse om gangen
	}
}

func handleConnection(c net.Conn) {

	//Les fra konneksjonen
	buffer := make([]byte, 1024)
	c.Read(buffer)

	//Skriv til konneksjonen
	c.Write([]byte("Hello from server\n"))

	// Steng forbindelse når det er klart.
	c.Close()
}




