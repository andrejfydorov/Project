package service

import (
	"Project/internal/repo"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (s *Service) Create(writer http.ResponseWriter, request *http.Request) {
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

		_, ok := s.repo.Cities[c.Id]
		if ok {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("Город с id %d уже существует.\n", c.Id)))
			return
		}
		s.repo.Cities[c.Id] = &c

		writer.WriteHeader(http.StatusCreated)
		writer.Write([]byte(fmt.Sprintf("Город %s успешно был создан\n", c.Name)))
		return
	}
	writer.WriteHeader(http.StatusBadRequest)
	return
}
