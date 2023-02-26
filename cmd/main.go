package main

import (
	"Project/internal/repo"
	"Project/internal/service"
	"log"
	"os"
	"os/signal"
)

func main() {

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	r := repo.New()
	defer r.Close()

	service.Start(r)

	for {
		select {
		case <-done:
			log.Println("Closing")
			r.Close()
			return
		}
	}

}
