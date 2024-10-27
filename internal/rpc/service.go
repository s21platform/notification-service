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

func (s *Service) GetNotification(ctx context.Context, in *notificationproto.NotificationIn) (*notificationproto.NotificationOut, error) {
	userUuid := ctx.Value(config.KeyUUID).(string)
	notifications, err := s.dbR.GetNotifications(ctx, userUuid, in.Limit, in.Offset)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Intenal Error: %v", err.Error())
	}
	var result []*notificationproto.Notification
	for _, notification := range notifications {
		result = append(result, &notificationproto.Notification{
			Id:     notification.Id,
			Text:   notification.Text,
			IsRead: notification.IsRead,
		})
	}
	return &notificationproto.NotificationOut{
		Notifications: result,
	}, nil
}
