package service

import (
	"Project/internal/repo"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func Info(wg *sync.WaitGroup, mutex *sync.Mutex, _repo *repo.Repo) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		go func() {
			wg.Add(1)
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()

			id := chi.URLParam(request, "id")
			//fmt.Println(id + "\n")
			idInt, err := strconv.Atoi(id)
			if err != nil {
				log.Fatalln(err)
			}

			city, ok := _repo.Cities[idInt]
			if ok {
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
			writer.WriteHeader(http.StatusBadRequest)
		}()
	}
}
