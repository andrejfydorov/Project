package repo

type City struct {
	Id         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Region     string `json:"region,omitempty"`
	District   string `json:"district,omitempty"`
	Population int    `json:"population,omitempty"`
	Foundation int    `json:"foundation,omitempty"`
}

type Repo struct {
	сities map[int]*City
}

func (r *Repo) Add(city *City) bool {
	_, ok := r.сities[city.Id]
	if ok {
		return false
	} else {
		r.сities[city.Id] = city
		return true
	}
}

func (r *Repo) IsExist(id int) bool {
	_, ok := r.сities[id]
	if ok {
		return true
	} else {
		return false
	}
}

func (r *Repo) Get(id int) *City {
	city, ok := r.сities[id]
	if ok {
		return city
	} else {
		return nil
	}
}

func (r *Repo) GetAll() []*City {
	var сities []*City
	for _, city := range r.сities {
		сities = append(сities, city)
	}
	return сities
}

func (r *Repo) Update(city *City) bool {
	city, ok := r.сities[city.Id]
	if ok {
		r.сities[city.Id] = city
		return true
	} else {
		return false
	}
}
