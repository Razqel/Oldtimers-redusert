package main

import ("fmt"
		"os"
		"os/signal"
		"syscall"
		"time"
)


func main() {
	for {
		fmt.Println("Hei folkens se på meg. CTRL-C FOR STOP")
		time.Sleep(time.Second * 5)

		var gracefulStop = make(chan os.Signal)
		signal.Notify(gracefulStop, syscall.SIGTERM)
		signal.Notify(gracefulStop, syscall.SIGINT)
		go func() {
			sig := <-gracefulStop
			fmt.Printf("Fanget opp signal!: %+v", sig)
			fmt.Println(" Vent 2 sekunder for avslutting av prosesser!")
			time.Sleep(2*time.Second)
			os.Exit(0)
		}()
	}
}

