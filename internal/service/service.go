package service

import (
	"Project/internal/repo"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func Start(r *repo.Repo) {

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/cities/{id:[0-9]+}", Info(r))
	router.Get("/cities", GetAll(r))
	router.Post("/cities", Create(r))
	router.Delete("/cities/{id:[0-9]+}", Delete(r))
	router.Put("/cities/{id:[0-9]+}", Update(r))
	router.Get("/cities/region", InfoRegion(r))
	router.Get("/cities/district", InfoDistrict(r))
	router.Get("/cities/population", InfoPopulation(r))
	router.Get("/cities/foundation", InfoFoundation(r))

	go func() {
		err := http.ListenAndServe(":8080", router)
		if err != nil {
			log.Fatalln(err)
		}
	}()

}
