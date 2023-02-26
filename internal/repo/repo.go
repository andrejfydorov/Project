package repo

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type City struct {
	Id         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Region     string `json:"region,omitempty"`
	District   string `json:"district,omitempty"`
	Population int    `json:"population,omitempty"`
	Foundation int    `json:"foundation,omitempty"`
}

type Repo struct {
	mutex  sync.Mutex
	сities map[int]*City
}

func New() *Repo {
	repo := Repo{
		сities: map[int]*City{},
	}

	repo.Open()

	return &repo
}

func (r *Repo) Add(city *City) bool {

	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, ok := r.сities[city.Id]
	if ok {
		return false
	} else {
		r.сities[city.Id] = city
		return true
	}
}

func (r *Repo) Update(city *City) bool {

	r.mutex.Lock()
	defer r.mutex.Unlock()

	city, ok := r.сities[city.Id]
	if ok {
		r.сities[city.Id] = city
		return true
	} else {
		return false
	}
}

func (r *Repo) Delete(id int) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, ok := r.сities[id]
	if ok {
		delete(r.сities, id)
		return true
	}
	return false
}

func (r *Repo) IsExist(id int) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, ok := r.сities[id]
	if ok {
		return true
	} else {
		return false
	}
}

func (r *Repo) Get(id int) *City {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	city, ok := r.сities[id]
	if ok {
		return city
	} else {
		return nil
	}
}

func (r *Repo) GetAll() []*City {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	сities := []*City{}
	for _, city := range r.сities {
		сities = append(сities, city)
	}

	return сities
}

func (r *Repo) Open() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	file, err := os.Open("resources/cities.csv")
	if err != nil {
		log.Println("Unable to open file:", err)
		log.Fatalln(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalln(err)
			}
		}

		str := strings.Split(line, ",")

		var c City
		key, err := strconv.Atoi(str[0])
		if err != nil {
			log.Fatalln(err)
		}
		c.Id = key
		c.Name = str[1]
		c.Region = str[2]
		c.District = str[3]
		i, err := strconv.Atoi(str[4])
		if err != nil {
			log.Fatalln(err)
		}
		c.Population = i
		str[5] = strings.Trim(str[5], "\r\n")
		i, err = strconv.Atoi(str[5])
		if err != nil {
			log.Fatalln(err)
		}
		c.Foundation = i

		r.сities[key] = &c
	}
}

func (r *Repo) Close() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	file, err := os.Create("resources/cities_.csv")
	if err != nil {
		log.Println("Unable to open file:", err)
		log.Fatalln(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, city := range r.сities {
		cityFull := fmt.Sprintf("%d,%s,%s,%s,%d,%d", city.Id, city.Name, city.Region, city.District, city.Population, city.Foundation)

		_, err := writer.WriteString(cityFull)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = writer.WriteString("\n")
		if err != nil {
			log.Fatalln(err)
		}

	}

	writer.Flush()
}
