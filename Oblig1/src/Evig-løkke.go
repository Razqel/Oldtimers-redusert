package main

import ("fmt"
"os"
	"os/signal"
	"log"
)

func main() {
 for i := 0 ;  ; i++ {
 	fmt.Println(i)
	 c := make(chan os.Signal, 1)
	 signal.Notify(c, os.Interrupt)
	 go func() {
		 for sig := range c {
			 log.Printf("captured %v, stopping Evig-løkke", sig)

			 os.Exit(1)
		 }
	 }()

 }

}
