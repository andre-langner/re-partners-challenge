package calculatepacksservice

import (
	"fmt"
	"slices"
)

type Repo interface {
	FindAvailablePacks() []int
}

type Service struct {
	repo Repo
}

func New(repo Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CalculatePacks(itemsOrdered int) map[int]int {
	availablePacks := s.repo.FindAvailablePacks()
	slices.Reverse(availablePacks)

	fmt.Printf("%v\n", availablePacks)

	packs := make([]int, 0)
	unpackedItems := itemsOrdered

	// calculate packs
	for _, packSize := range availablePacks {
		packsCount := unpackedItems / packSize

		for j := 0; j < packsCount; j++ {
			unpackedItems -= packSize
			packs = append(packs, packSize)
		}
	}

	// add the rest of the items
	if unpackedItems > 0 {
		slices.Sort(availablePacks)
		for _, packSize := range availablePacks {
			if packSize > unpackedItems {
				packs = append(packs, packSize)
				break
			}
		}
	}

	// summarize
	calculatedPacks := map[int]int{}
	for _, pack := range packs {
		calculatedPacks[pack]++
	}

	// merge packs to avoid breaking rule 3
	for calculatedPackSize, count := range calculatedPacks {
		if count == 2 {
			for _, packSize := range availablePacks {
				if calculatedPackSize < packSize {
					delete(calculatedPacks, calculatedPackSize)
					calculatedPacks[packSize]++
					break
				}
			}
		}
	}

	return calculatedPacks
}
