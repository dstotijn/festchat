package chat

import "errors"

var errMemberNotFound = errors.New("member not found")

// Member represents a person with a name and unique phone number.
type Member struct {
	Name        string
	PhoneNumber string
}
