package calculatepackshandler_test

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/suite"

	"re_partners/internal/handlers/calculatepackshandler"
	"re_partners/internal/repository/packsrepo"
	"re_partners/internal/services/calculatepacksservice"
)

type ITSuite struct {
	suite.Suite

	Handler *calculatepackshandler.Handler
}

func (s *ITSuite) SetupSuite() {
	repo := packsrepo.New()

	service := calculatepacksservice.New(repo)

	s.Handler = calculatepackshandler.New(service)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ITSuite))
}

func (s *ITSuite) TestInvalidOrderItem() {
	apitest.New().
		Handler(s.Handler).
		Get("/packs").
		QueryParams(map[string]string{
			"orderItems": "abc",
		}).
		Expect(s.T()).
		Header("Content-Type", "application/json").
		Status(http.StatusBadRequest).
		Body(`{"message":"schema: error converting value for \"orderItems\""}`).
		End()
}

func (s *ITSuite) Test_1_Item() {
	apitest.New().
		Handler(s.Handler).
		Get("/packs").
		QueryParams(map[string]string{
			"orderItems": "1",
		}).
		Expect(s.T()).
		Header("Content-Type", "application/json").
		Status(http.StatusOK).
		Body(`{"packs":[{"pack":250,"quantity":1}]}`).
		End()
}

func (s *ITSuite) Test_250_Items() {
	apitest.New().
		Handler(s.Handler).
		Get("/packs").
		QueryParams(map[string]string{
			"orderItems": "250",
		}).
		Expect(s.T()).
		Header("Content-Type", "application/json").
		Status(http.StatusOK).
		Body(`{"packs":[{"pack":250,"quantity":1}]}`).
		End()
}

func (s *ITSuite) Test_251_Items() {
	apitest.New().
		Handler(s.Handler).
		Get("/packs").
		QueryParams(map[string]string{
			"orderItems": "251",
		}).
		Expect(s.T()).
		Header("Content-Type", "application/json").
		Status(http.StatusOK).
		Body(`{"packs":[{"pack":500,"quantity":1}]}`).
		End()
}

func (s *ITSuite) Test_501_Items() {
	apitest.New().
		Handler(s.Handler).
		Get("/packs").
		QueryParams(map[string]string{
			"orderItems": "501",
		}).
		Expect(s.T()).
		Header("Content-Type", "application/json").
		Status(http.StatusOK).
		Body(`{"packs":[{"pack":250,"quantity":1},{"pack":500,"quantity":1}]}`).
		End()
}

func (s *ITSuite) Test_12001_Items() {
	apitest.New().
		Handler(s.Handler).
		Get("/packs").
		QueryParams(map[string]string{
			"orderItems": "12001",
		}).
		Expect(s.T()).
		Header("Content-Type", "application/json").
		Status(http.StatusOK).
		Body(`{"packs":[{"pack":250,"quantity":1},{"pack":2000,"quantity":1},{"pack":5000,"quantity":2}]}`).
		End()
}
