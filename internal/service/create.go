package service

import (
	"Project/internal/repo"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func Create(wg *sync.WaitGroup, mutex *sync.Mutex, _repo *repo.Repo) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		go func() {
			wg.Add(1)
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()

			if request.Method == "POST" {
				content, err := ioutil.ReadAll(request.Body)
				if err != nil {
					writer.WriteHeader(http.StatusInternalServerError)
					writer.Write([]byte(err.Error()))
					return
				}
				defer request.Body.Close()

				var c repo.City
				if err := json.Unmarshal(content, &c); err != nil {
					writer.WriteHeader(http.StatusInternalServerError)
					writer.Write([]byte(err.Error()))
					fmt.Println(err)
					return
				}

				_, ok := _repo.Cities[c.Id]
				if ok {
					writer.WriteHeader(http.StatusBadRequest)
					writer.Write([]byte(fmt.Sprintf("Город с id %d уже существует.\n", c.Id)))
					return
				}
				_repo.Cities[c.Id] = &c

				writer.WriteHeader(http.StatusCreated)
				writer.Write([]byte(fmt.Sprintf("Город %s успешно был создан\n", c.Name)))
				return
			}
			writer.WriteHeader(http.StatusBadRequest)
			return
		}()
	}

}
