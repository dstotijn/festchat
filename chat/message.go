package chat

// ReceivedMessage represents a received text message.
type ReceivedMessage struct {
	Originator string
	Recipient  string
	Payload    string
}

type relayMessageRequest struct {
	Payload    string
	Originator Member
	Group      Group
}
