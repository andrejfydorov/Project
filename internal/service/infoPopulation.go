package service

import (
	"Project/internal/repo"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *Service) InfoPopulation(writer http.ResponseWriter, request *http.Request) {
	content, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	defer request.Body.Close()

	type Population struct {
		Start int `json:"start,omitempty"`
		Stop  int `json:"stop,omitempty"`
	}
	p := Population{}

	if err := json.Unmarshal(content, &p); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		fmt.Println(err)
		return
	}

	var cities []*repo.City
	for _, city := range s.repo.Cities {
		if city.Population >= p.Start && city.Population <= p.Stop {
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
		writer.Write([]byte(fmt.Sprintf("Города в указанном диапазоне %d - %d не найдены.\n", p.Start, p.Stop)))
		return
	}

	writer.WriteHeader(http.StatusBadRequest)
	return
}

