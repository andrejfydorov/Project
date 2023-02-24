package repo

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func Close(wg *sync.WaitGroup, mutex *sync.Mutex, _repo *Repo) {
	wg.Add(1)
	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()

	file, err := os.Create("resources/cities_.csv")
	if err != nil {
		log.Println("Unable to open file:", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, city := range _repo.сities {
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
