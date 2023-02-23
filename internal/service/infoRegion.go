package service

import (
	"Project/internal/repo"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func (s *Service) InfoRegion(writer http.ResponseWriter, request *http.Request) {
	content, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	defer request.Body.Close()

	type Region struct {
		Region string `json:"region,omitempty"`
	}
	var r Region

	if err := json.Unmarshal(content, &r); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		fmt.Println(err)
		return
	}

	var cities []*repo.City
	for _, city := range s.repo.Cities {
		if strings.ToUpper(city.Region) == strings.ToUpper(r.Region) {
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
		writer.Write([]byte(fmt.Sprintf("Города в регионе %s не найдены.\n", r.Region)))
		return
	}
	writer.WriteHeader(http.StatusBadRequest)
	return
}
