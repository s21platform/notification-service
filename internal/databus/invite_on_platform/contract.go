package invite_on_platform

import "context"

type InviteMailClient interface {
	SendEmail(subject string, to string, content string) error
}

type UserClient interface {
	GetLoginByUuid(ctx context.Context, uuid string) (string, error)
}
