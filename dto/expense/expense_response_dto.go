package expense

import (
	"time"
	"uuid"
)

type ResponseDto struct {
	Id        uuid.UUID `json:"id" db:"id"`
	Amount    float64   `json:"amount" db:"amount"`
	Category  string    `json:"category" db:"category"`
	Note      *string   `json:"note" db:"note"`
	SpentOn   time.Time `json:"spent_on" db:"spent_on"`
	CreatedOn time.Time `json:"created_on" db:"created_on"`
}
