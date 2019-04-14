package chat

// Group represents a group of people that chat.
type Group struct {
	PhoneNumber string
	Members     []Member
}

// recipientPhoneNumbers returns the phone numbers of the group members,
// with the member passed in this function excluded.
func (group Group) recipientPhoneNumbers(m Member) []string {
	recipients := make([]string, len(group.Members)-1)

	for _, member := range group.Members {
		if member != m {
			recipients = append(recipients, member.PhoneNumber)
		}
	}

	return recipients
}

func (group Group) findMemberByPhoneNumber(phoneNum string) (*Member, error) {
	for i, member := range group.Members {
		if member.PhoneNumber == phoneNum {
			return &group.Members[i], nil
		}
	}

	return nil, errMemberNotFound
}
