package app

import (
	"Calorie_calculator/internal/pkg/utils"
	"errors"
	"net/http"
)

func (s Service) CalcProdFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.calcProd(w, r)
	default:
		utils.WriteInternalError(w, errors.New("unknown type for DriverFunc"), "")
	}
}
