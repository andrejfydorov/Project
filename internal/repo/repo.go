package repo

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
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
	Cities map[int]*City
}

func New() (*Repo, error) {

	repo := Repo{Cities: map[int]*City{}}

	file, err := os.Open("resources/cities.csv")
	if err != nil {
		log.Println("Unable to open file:", err)
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Println(err)
			}
		}

		str := strings.Split(line, ",")

		//fmt.Print(line)

		var c City
		key, err := strconv.Atoi(str[0])
		if err != nil {
			log.Println(err)
		}
		c.Id = key
		c.Name = str[1]
		c.Region = str[2]
		c.District = str[3]
		i, err := strconv.Atoi(str[4])
		if err != nil {
			log.Println(err)
		}
		c.Population = i
		str[5] = strings.Trim(str[5], "\r\n")
		i, err = strconv.Atoi(str[5])
		if err != nil {
			log.Println(err)
		}
		c.Foundation = i

		repo.Cities[key] = &c
	}

	//for i, i2 := range repo.Cities {
	//	fmt.Println(i, i2)
	//}

	return &repo, nil
}

func (r *Repo) Close() {

	file, err := os.Create("resources/cities_.csv")
	if err != nil {
		log.Println("Unable to open file:", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, city := range r.Cities {
		cityFull := fmt.Sprintf("%d,%s,%s,%s,%d,%d", city.Id, city.Name, city.Region, city.District, city.Population, city.Foundation)

		_, err := writer.WriteString(cityFull)
		if err != nil {
			log.Println(err)
		}
		_, err = writer.WriteString("\n")
		if err != nil {
			log.Println(err)
		}

	}
	writer.Flush()
}

func (r *Repo) Info(id int) City {
	return *r.Cities[id]
}

func (r *Repo) Delete(id int) bool {
	_, ok := r.Cities[id]
	if ok {
		delete(r.Cities, id)
		return true
	}
	return false
}
