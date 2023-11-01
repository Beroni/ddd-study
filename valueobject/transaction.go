package valueobject

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct { 
	from uuid.UUID
	to uuid.UUID
	amount float64
	createdAt time.Time
}