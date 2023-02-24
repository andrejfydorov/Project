package repo

func (r *Repo) Delete(id int) bool {
	_, ok := r.сities[id]
	if ok {
		delete(r.сities, id)
		return true
	}
	return false
}
