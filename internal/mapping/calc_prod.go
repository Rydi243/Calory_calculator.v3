package mapping

import (
	"Calorie_calculator/internal/contract"
	"Calorie_calculator/postgres/public/model"
)

func CreateToDB(req *contract.CreateProdRequest) *model.Calorie {
	return &model.Calorie{
		Product: req.Product,
		Weight:  int32(req.Weight),
		Calorie: int32(req.Calorie),
	}
}

func CreateToResp(m *model.Calorie) *contract.CreateInfo {
	return &contract.CreateInfo{
		ID:      int(m.ID),
		Product: m.Product,
		Weight:  int(m.Weight),
		Calorie: int(m.Calorie),
	}
}

func CalorieToResp(models []*model.Calorie) *contract.CaloriInfo {
	items := make([]*contract.CreateInfo, 0, len(models))

	for _, m := range models {
		items = append(items, CreateToResp(m))
	}

	return &contract.CaloriInfo{Items: items}
}

func CalcToDB(r *contract.CalcProdRequest, i int) *model.Calorie {
	return &model.Calorie{
		Product: r.Product[i],
		Weight:  int32(r.Weight[i]),
	}
}

func CalcToResp(c float64) *contract.CalcResp {
	return &contract.CalcResp{
		Calc: c,
	}
}
