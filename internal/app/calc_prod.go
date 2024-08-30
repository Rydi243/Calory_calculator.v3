package app

import (
	"Calorie_calculator/internal/contract"
	"Calorie_calculator/internal/mapping"
	"Calorie_calculator/internal/pkg/utils"
	"encoding/json"
	"fmt"
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

	fmt.Println(calc.Product[0])

	for i, _ := range calc.Product {
		mapp := mapping.CalcToDB(&calc, i)
		cal, err = s.store.Calc(mapp)
		if err != nil {
			utils.WriteInternalError(w, err, "store.Calc")
			return
		}

		resit += cal
	}
	fmt.Println(resit)

	// res1 := mapping.CalcToDB(&calc, 0)
	// res2 := mapping.CalcToDB(&calc, 1)
	// fmt.Println(res1)
	// fmt.Println(res2)

	// result, err := s.store.Calc(res1, res2)
	// if err != nil {
	// 	utils.WriteInternalError(w, err, "store.Calc")
	// }
	// fmt.Println(result)

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
