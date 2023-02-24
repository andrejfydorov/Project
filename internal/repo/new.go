package repo

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func New() *Repo {

	repo := Repo{Cities: map[int]*City{}}

	file, err := os.Open("resources/cities.csv")
	if err != nil {
		log.Println("Unable to open file:", err)
		log.Println(err)
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

	return &repo
}
