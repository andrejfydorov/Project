package service

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func (s *Service) Delete(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	//fmt.Println(id + "\n")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalln(err)
	}

	if s.repo.Delete(idInt) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(fmt.Sprintf("Город с id %d удален.\n", idInt)))
		return
	} else {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(fmt.Sprintf("Город с id %d не найден.\n", idInt)))
		return
	}
	writer.WriteHeader(http.StatusBadRequest)
	return
}
