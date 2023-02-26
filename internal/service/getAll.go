package service

import (
	"Project/internal/repo"
	"encoding/json"
	"log"
	"net/http"
)

func GetAll(r *repo.Repo) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		cities := r.GetAll()
		js, err := json.Marshal(cities)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
			log.Fatalln(err)
			return
		}

		writer.WriteHeader(http.StatusOK)
		writer.Write(js)
	}
}
