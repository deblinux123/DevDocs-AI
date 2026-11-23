package document

import "github.com/google/uuid"

func NewID() string {

	return uuid.NewString()
}
