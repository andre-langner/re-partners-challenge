package packsrepo

import "slices"

var data = []int{250, 500, 1000, 2000, 5000}

type Repository struct{}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) FindAvailablePacks() []int {
	slices.Sort(data)
	return data
}
