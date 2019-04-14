package sms

// Sender represents an interface for sending SMS messages.
type Sender interface {
	Send(msg Message) error
}
