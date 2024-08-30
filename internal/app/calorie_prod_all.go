package app

import (
	"Calorie_calculator/internal/mapping"
	"Calorie_calculator/internal/pkg/utils"
	"encoding/json"
	"net/http"
)

func (s Service) calorieProdAll(w http.ResponseWriter, r *http.Request) {
	resDB, err := s.store.GetAllProduct()
	if err != nil {
		utils.WriteInternalError(w, err, "store.GetAllProd")
	}

	res := mapping.CalorieToResp(resDB)
	resByte, err := json.Marshal(res)
	if err != nil {
		utils.WriteInternalError(w, err, "json.Marshal")
	}

	_, err = w.Write(resByte)
	if err != nil {
		utils.WriteInternalError(w, err, "w.Write")
	}

}
