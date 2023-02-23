package main

import (
	"Project/internal/repo"
	"Project/internal/service"
	"log"
	"os"
	"os/signal"
)

var r repo.Repo

func main() {

	done := make(chan os.Signal, 1)

	r, err := repo.New()
	if err != nil {
		log.Println(err)
		return
	}
	//defer r.Close()

	err = service.New(r)
	if err != nil {
		log.Println(err)
		return
	}

	go func() {
		for {
			select {
			case <-done:
				r.Close()
				return
			default:
				signal.Notify(done, os.Interrupt)
			}
		}

	}()

}
