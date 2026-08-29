package expense

type CreateDto struct {
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
	Note     *string `json:"note"`
	SpentOn  string  `json:"spent_on"`
}
