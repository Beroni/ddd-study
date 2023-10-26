package entity

import (
	"github.com/google/uuid"
)

// A estrutura Item representa um item em todos os dominios do projeto.
type Item struct {
	ID uuid.UUID
	Name string
	Price float64
	Quantity int64
}