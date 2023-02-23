package main

import (
	"Project/internal/repo"
	"Project/internal/service"
	"log"
)

func main() {

	r, err := repo.New()
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Close()

	err = service.New(r)
	if err != nil {
		log.Println(err)
		return
	}

}
