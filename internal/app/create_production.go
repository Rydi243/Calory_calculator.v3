package app

import (
	"Calorie_calculator/internal/contract"
	"Calorie_calculator/internal/mapping"
	"Calorie_calculator/internal/pkg/utils"
	"encoding/json"
	"net/http"
)

func (s Service) createProd(w http.ResponseWriter, r *http.Request) {
	var create contract.CreateProdRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&create)
	if err != nil {
		utils.WriteInternalError(w, err, "decoder.Decode")
		//return
	}

	createModel := mapping.CreateToDB(&create)
	resDB, err := s.store.InsertProduct(createModel)
	if err != nil {
		utils.WriteInternalError(w, err, "store.InsertProduct")
		//return
	}

	res := mapping.CreateToResp(resDB)
	resByte, err := json.Marshal(res)
	if err != nil {
		utils.WriteInternalError(w, err, "json.Marshal")
		//return
	}

	_, err = w.Write(resByte)
	if err != nil {
		utils.WriteInternalError(w, err, "w.Write")
		//return
	}
}
