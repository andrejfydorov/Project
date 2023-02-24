package service

import (
	"Project/internal/repo"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func Update(wg *sync.WaitGroup, mutex *sync.Mutex, _repo *repo.Repo) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		go func() {
			wg.Add(1)
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()

			content, err := ioutil.ReadAll(request.Body)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				writer.Write([]byte(err.Error()))
				return
			}
			defer request.Body.Close()

			id := chi.URLParam(request, "id")
			fmt.Println(id + "\n")
			idInt, err := strconv.Atoi(id)
			if err != nil {
				log.Fatalln(err)
			}

			city, ok := _repo.Cities[idInt]
			if ok {
				type Population struct {
					Population int `json:"population,omitempty"`
				}
				var p Population

				if err := json.Unmarshal(content, &p); err != nil {
					writer.WriteHeader(http.StatusInternalServerError)
					writer.Write([]byte(err.Error()))
					fmt.Println(err)
					return
				}

				_repo.Cities[idInt].Population = p.Population

				writer.WriteHeader(http.StatusOK)
				writer.Write([]byte(fmt.Sprintf("Город %s успешно обновлен.\n", city.Name)))
				return
			} else {
				writer.WriteHeader(http.StatusBadRequest)
				writer.Write([]byte(fmt.Sprintf("Город с %d не найден.\n", idInt)))
				return
			}

			writer.WriteHeader(http.StatusBadRequest)
		}()
	}
}
