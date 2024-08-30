package app

import (
	"Calorie_calculator/internal/contract"
	"Calorie_calculator/internal/mapping"
	"Calorie_calculator/internal/pkg/utils"
	"encoding/json"
	"net/http"
)

func (s Service) calcProd(w http.ResponseWriter, r *http.Request) {
	var calc contract.CalcProdRequest
	var resit, cal float64

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&calc)
	if err != nil {
		utils.WriteInternalError(w, err, "decoder.Decode")
	}

	for i := range calc.Product {
		mapp := mapping.CalcToDB(&calc, i)
		cal, err = s.store.Calc(mapp)
		if err != nil {
			utils.WriteInternalError(w, err, "store.Calc")
			return
		}

		resit += cal
	}

	res := mapping.CalcToResp(resit)
	resByte, err := json.Marshal(res)
	if err != nil {
		utils.WriteInternalError(w, err, "json.Marshal")
	}

	_, err = w.Write(resByte)
	if err != nil {
		utils.WriteInternalError(w, err, "w.Write")
	}
}
