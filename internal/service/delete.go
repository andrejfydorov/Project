package service

import (
	"Project/internal/repo"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func Delete(wg *sync.WaitGroup, mutex *sync.Mutex, _repo *repo.Repo) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		go func() {
			wg.Add(1)
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()

			id := chi.URLParam(request, "id")

			idInt, err := strconv.Atoi(id)
			if err != nil {
				log.Fatalln(err)
			}

			if _repo.Delete(idInt) {
				writer.WriteHeader(http.StatusOK)
				writer.Write([]byte(fmt.Sprintf("Город с id %d удален.\n", idInt)))
				return
			} else {
				writer.WriteHeader(http.StatusBadRequest)
				writer.Write([]byte(fmt.Sprintf("Город с id %d не найден.\n", idInt)))
				return
			}
		}()
	}
}
