package service

import (
	"Project/internal/repo"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"sync"
)

func Start(wg *sync.WaitGroup, mutex *sync.Mutex, _repo *repo.Repo) {

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/info/{id:[0-9]+}", Info(wg, mutex, _repo))
	router.Post("/create", Create(wg, mutex, _repo))
	router.Delete("/delete/{id:[0-9]+}", Delete(wg, mutex, _repo))
	router.Put("/update/{id:[0-9]+}", Update(wg, mutex, _repo))
	router.Post("/info/region", InfoRegion(wg, mutex, _repo))
	router.Get("/info/district", InfoDistrict(wg, mutex, _repo))
	router.Post("/info/population", InfoPopulation(wg, mutex, _repo))
	router.Post("/info/foundation", InfoFoundation(wg, mutex, _repo))

	go func() {
		err := http.ListenAndServe(":8080", router)
		if err != nil {
			log.Fatalln(err)
		}
	}()

}
