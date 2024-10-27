package rpc

import "context"

type DbRepo interface {
	GetCountNotification(ctx context.Context, userUuid string) (int64, error)
}
