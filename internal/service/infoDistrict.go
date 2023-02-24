package service

import (
	"Project/internal/repo"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

func InfoDistrict(wg *sync.WaitGroup, mutex *sync.Mutex, _repo *repo.Repo) http.HandlerFunc {
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
				log.Fatalln(err)
				return
			}

			defer request.Body.Close()

			type District struct {
				District string `json:"district,omitempty"`
			}

			var d District

			if err := json.Unmarshal(content, &d); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				writer.Write([]byte(err.Error()))
				log.Fatalln(err)
				return
			}

			var cities []*repo.City
			for _, city := range _repo.GetAll() {
				if strings.ToUpper(city.District) == strings.ToUpper(d.District) {
					cities = append(cities, city)
				}
			}

			if len(cities) > 0 {
				js, err := json.Marshal(cities)
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
				writer.Write([]byte(fmt.Sprintf("Города в округе %s не найдены.\n", d.District)))
				return
			}
		}()
	}
}
