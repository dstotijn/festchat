package memory

import (
	"errors"

	"github.com/dstotijn/festchat/chat"
)

// ErrGroupNotFound is used when a group was not found.
var ErrGroupNotFound = errors.New("memory: group not found")

// Store represents an in memory store of groups.
type Store struct {
	groups map[string]*chat.Group
}

// NewStore returns a new Store.
func NewStore() *Store {
	return &Store{
		groups: map[string]*chat.Group{
			// TODO: Insert fixture data?
		},
	}
}

// FindGroupByPhoneNumber finds a group by its phone number in the store.
func (s Store) FindGroupByPhoneNumber(phoneNum string) (*chat.Group, error) {
	group, ok := s.groups[phoneNum]
	if !ok {
		return nil, ErrGroupNotFound
	}

	return group, nil
}
