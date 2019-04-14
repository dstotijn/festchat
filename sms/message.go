package sms

// Message represents an SMS message.
type Message struct {
	Originator string
	Recipients []string
	Payload    string
}
