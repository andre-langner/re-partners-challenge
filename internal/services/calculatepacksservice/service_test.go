package calculatepacksservice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"re_partners/internal/services/calculatepacksservice"
	"re_partners/internal/services/calculatepacksservice/mocks"
)

func TestService_CalculatePacks(t *testing.T) {
	cases := map[string]struct {
		itemsOrdered int
		expected     map[int]int
	}{
		"1-item": {
			itemsOrdered: 1,
			expected:     map[int]int{250: 1},
		},
		"250-items": {
			itemsOrdered: 250,
			expected:     map[int]int{250: 1},
		},
		"251-items": {
			itemsOrdered: 251,
			expected:     map[int]int{500: 1},
		},
		"501-items": {
			itemsOrdered: 501,
			expected:     map[int]int{500: 1, 250: 1},
		},
		"12001-items": {
			itemsOrdered: 12001,
			expected:     map[int]int{5000: 2, 2000: 1, 250: 1},
		},
	}

	for title, tc := range cases {
		t.Run(title, func(t *testing.T) {
			repo := mocks.NewRepo(t)
			repo.On("FindAvailablePacks").Return([]int{250, 500, 1000, 2000, 5000})

			service := calculatepacksservice.New(repo)

			calculatedPacks := service.CalculatePacks(tc.itemsOrdered)

			assert.Equal(t, tc.expected, calculatedPacks)
		})
	}
}
