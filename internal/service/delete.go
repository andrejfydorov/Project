package service

import (
	"Project/internal/repo"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func Delete(r *repo.Repo) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id := chi.URLParam(request, "id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Fatalln(err)
		}

		if r.Delete(idInt) {
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte(fmt.Sprintf("Город с id %d удален.\n", idInt)))
			return
		} else {
			writer.WriteHeader(http.StatusNotFound)
			writer.Write([]byte(fmt.Sprintf("Город с id %d не найден.\n", idInt)))
			return
		}
	}
}
