package user_invite

type InviteMailClient interface {
	SendEmail(subject string, to string, content string) error
}
