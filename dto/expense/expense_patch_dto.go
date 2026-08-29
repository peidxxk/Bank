package expense

import "uuid"

type PatchDto struct {
	Id       uuid.UUID `json:"id" db:"id"`
	Amount   float64   `json:"amount" db:"amount"`
	Category string    `json:"category" db:"category"`
	Note     *string   `json:"note" db:"note"`
}
