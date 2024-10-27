package rpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	notificationproto "github.com/s21platform/notification-proto/notification-proto"

	"notification-service/internal/config"
)

type Service struct {
	notificationproto.UnimplementedNotificationServiceServer
	dbR DbRepo
}

func New(dbR DbRepo) *Service {
	return &Service{dbR: dbR}
}

func (s *Service) GetNotificationCount(ctx context.Context, _ *notificationproto.Empty) (*notificationproto.NotificationCountOut, error) {
	userUuid := ctx.Value(config.KeyUUID).(string)
	count, err := s.dbR.GetCountNotification(ctx, userUuid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Intenal Error: %v", err.Error())
	}
	return &notificationproto.NotificationCountOut{
		Count: count,
	}, nil
}
