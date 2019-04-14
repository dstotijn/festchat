package chat

import (
	"fmt"

	"github.com/dstotijn/festchat/sms"
)

// Service is used for relaying chat messages amongst group members.
type Service struct {
	smsSender  sms.Sender
	repository Repository
}

// NewService returns a new Service.
func NewService(smsSender sms.Sender, repository Repository) *Service {
	return &Service{
		smsSender:  smsSender,
		repository: repository,
	}
}

// RelayMessage relays a received message to related group participants.
func (svc Service) RelayMessage(msg ReceivedMessage) error {
	group, err := svc.repository.FindGroupByPhoneNumber(msg.Recipient)
	if err != nil {
		return err
	}

	originator, err := group.findMemberByPhoneNumber(msg.Originator)
	if err != nil {
		return fmt.Errorf("chat: could not find originator: %v", err)
	}

	req := relayMessageRequest{
		Payload:    msg.Payload,
		Originator: *originator,
		Group:      *group,
	}

	relayMsg := sms.Message{
		Originator: req.Group.PhoneNumber,
		Recipients: req.Group.recipientPhoneNumbers(req.Originator),
		Payload:    fmt.Sprintf("%v: %v", req.Originator.Name, req.Payload),
	}

	if err := svc.smsSender.Send(relayMsg); err != nil {
		return fmt.Errorf("chat: could not send SMS: %v", err)
	}

	return nil
}
