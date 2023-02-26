package service

import (
	"Project/internal/repo"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Create(r *repo.Repo) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		content, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
			log.Fatalln(err)
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

		ok := r.IsExist(c.Id)
		if ok {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("Город с id %d уже существует.\n", c.Id)))
			return
		}

		r.Add(&c)

		writer.Header().Set("Location", fmt.Sprintf("/cities/%d", c.Id))
		writer.WriteHeader(http.StatusCreated)
		writer.Write([]byte(fmt.Sprintf("Город %s успешно был создан\n", c.Name)))
		return
	}
}
