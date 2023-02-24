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
	Cities map[int]*City
}

//func (r *Repo) Info(id int) City {
//	return *r.Cities[id]
//}
