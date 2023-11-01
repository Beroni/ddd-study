package entity

import (
	"github.com/google/uuid"
)

// A estrutura Person representa uma pessoa em todos os dominios do projeto.
type Person struct {
	ID    uuid.UUID
	Email string
	Name  string
}
