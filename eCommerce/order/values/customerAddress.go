package values

import "github.com/google/uuid"

type CustomerAddress struct {
	CustomerId uuid.UUID
	Address    Address
}
