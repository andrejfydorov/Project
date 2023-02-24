package main

import (
	"Project/internal/repo"
	"Project/internal/service"
	"log"
	"os"
	"os/signal"
	"sync"
)

var Repo *repo.Repo

func main() {

	var wg sync.WaitGroup
	var mutex sync.Mutex

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	Repo = repo.New()
	defer repo.Close(&wg, &mutex, Repo)

	service.Start(&wg, &mutex, Repo)

	wg.Wait()

	for {
		select {
		case <-done:
			log.Println("Closing")
			repo.Close(&wg, &mutex, Repo)
			wg.Wait()
			return
		}
	}

}
