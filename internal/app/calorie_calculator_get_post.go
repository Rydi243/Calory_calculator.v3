package app

import (
	"Calorie_calculator/internal/pkg/utils"
	"errors"
	"net/http"
)

func (s Service) CalculatorFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.createProd(w, r)
	case http.MethodGet:
		s.calorieProdAll(w, r)
	default:
		utils.WriteInternalError(w, errors.New("unknown type for DriverFunc"), "")
	}
}
