//go:generate mockgen -destination=mock_contract_test.go -package=${GOPACKAGE} -source=contract.go
package service

import (
	"context"

	"github.com/s21platform/notification-service/internal/model"
)

type DbRepo interface {
	GetCountNotification(ctx context.Context, userUuid string) (int64, error)
	GetNotifications(ctx context.Context, userUuid string, limit int64, offset int64) ([]model.Notification, error)
	MarkNotificationsAsRead(ctx context.Context, userUuid string, notificationId []int64) error
}
