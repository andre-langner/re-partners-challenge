package calculatepackshandler

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/gorilla/schema"
)

type RequestParameters struct {
	OrderItems *int `schema:"orderItems,required"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type PackResponse struct {
	Pack     int `json:"pack,omitempty"`
	Quantity int `json:"quantity,omitempty"`
}

type SuccessResponse struct {
	Packs []PackResponse `json:"packs"`
}

type Service interface {
	CalculatePacks(itemsOrdered int) map[int]int
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (handler *Handler) GetHeaders() map[string]string {
	return map[string]string{}
}

func (handler *Handler) GetMethod() string {
	return http.MethodGet
}

func (handler *Handler) GetPath() string {
	return "/packs"
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	// parse and validate query params
	params := new(RequestParameters)
	if err := schema.NewDecoder().Decode(params, r.Form); err != nil {
		handler.writeErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// calculate
	packs := handler.service.CalculatePacks(*params.OrderItems)

	handler.writeSuccessResponse(w, packs)
}

func (handler *Handler) writeErrorResponse(w http.ResponseWriter, message string, httpCode int) {
	bodyJSON, _ := json.Marshal(ErrorResponse{
		Message: message,
	})

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	_, _ = w.Write(bodyJSON)
}

func (handler *Handler) writeSuccessResponse(w http.ResponseWriter, packs map[int]int) {
	var packResponse []PackResponse
	for pack, quantity := range packs {
		packResponse = append(packResponse, PackResponse{
			Pack:     pack,
			Quantity: quantity,
		})
	}

	sort.SliceStable(packResponse, func(i, j int) bool {
		return packResponse[i].Pack < packResponse[j].Pack
	})

	bodyJSON, _ := json.Marshal(SuccessResponse{
		Packs: packResponse,
	})

	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(bodyJSON)
}
