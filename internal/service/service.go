package service

import (
	"Project/internal/repo"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

type Service struct {
	repo *repo.Repo
}

func New(r *repo.Repo) error {

	s := Service{repo: r}

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/info/{id:[0-9]+}", s.Info)
	router.Post("/create", s.Create)
	router.Post("/delete/{id:[0-9]+}", s.Delete)
	router.Post("/update/{id:[0-9]+}", s.Update)
	router.Post("/info/region", s.InfoRegion)
	router.Get("/info/district", s.InfoDistrict)
	router.Post("/info/population", s.InfoPopulation)
	router.Post("/info/foundation", s.InfoFoundation)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
