package contract

type CreateProdRequest struct {
	Product string `json:"product"`
	Weight  int    `json:"weight"`
	Calorie int    `json:"calorie"`
}

type CreateInfo struct {
	ID      int    `json:"id"`
	Product string `json:"product"`
	Weight  int    `json:"weight"`
	Calorie int    `json:"calorie"`
}

type CaloriInfo struct {
	Items []*CreateInfo `json:"items"`
}

type CalcProdRequest struct {
	Product []string  `json:"product"`
	Weight  []float64 `json:"weight"`
}

type CalcResp struct {
	Calc float64 `json:"calc"`
}
