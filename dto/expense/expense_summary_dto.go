package expense

type SummaryDto struct {
	Category string `json:"category" db:"category"`
}
