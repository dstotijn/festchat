package chat

// Repository defines an interface for chat related storageo ops.
type Repository interface {
	FindGroupByPhoneNumber(phoneNumber string) (*Group, error)
}
