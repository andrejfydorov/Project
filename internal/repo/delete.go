package repo

func (r *Repo) Delete(id int) bool {
	_, ok := r.Cities[id]
	if ok {
		delete(r.Cities, id)
		return true
	}
	return false
}
