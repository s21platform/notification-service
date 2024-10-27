package rpc

import (
	"context"
	"notification-service/internal/model"
)

type DbRepo interface {
	GetCountNotification(ctx context.Context, userUuid string) (int64, error)
	GetNotifications(ctx context.Context, userUuid string, limit int64, offset int64) ([]model.Notification, error)
}
