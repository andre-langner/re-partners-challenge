package packsrepo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"re_partners/internal/repository/packsrepo"
)

func TestBaseRateRepository_FindBaseRateByLoanTermAndCreditScore(t *testing.T) {
	repo := packsrepo.New()
	packs := repo.FindAvailablePacks()

	expected := []int{250, 500, 1000, 2000, 5000}

	assert.ElementsMatch(t, packs, expected)
}
