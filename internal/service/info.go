package service

import (
	"Project/internal/repo"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func Info(r *repo.Repo) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		id := chi.URLParam(request, "id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Fatalln(err)
		}

		city := r.Get(idInt)
		if city != nil {
			js, err := json.Marshal(city)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				writer.Write([]byte(err.Error()))
				log.Fatalln(err)
				return
			}
			writer.WriteHeader(http.StatusOK)
			writer.Write(js)
			return
		} else {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("Город с id %d не найден.\n", idInt)))
			return
		}
	}
}
