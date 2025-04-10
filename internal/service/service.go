package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/s21platform/notification-service/pkg/notification"

	"github.com/s21platform/notification-service/internal/config"
)

var ErrNotificationNotFound = errors.New("notification not found or already read")

type Service struct {
	notification.UnimplementedNotificationServiceServer
	dbR DbRepo
}

func New(repo DbRepo) *Service {
	return &Service{
		dbR: repo,
	}
}

func (s *Service) GetNotificationCount(ctx context.Context, _ *emptypb.Empty) (*notification.NotificationCountOut, error) {
	log.Println("GetNotificationCount")
	userUuid := ctx.Value(config.KeyUUID).(string)
	count, err := s.dbR.GetCountNotification(ctx, userUuid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Intenal Error: %v", err.Error())
	}
	return &notification.NotificationCountOut{
		Count: count,
	}, nil
}

func (s *Service) GetNotification(ctx context.Context, in *notification.NotificationIn) (*notification.NotificationOut, error) {
	log.Println("GetNotification")
	userUuid := ctx.Value(config.KeyUUID).(string)
	notifications, err := s.dbR.GetNotifications(ctx, userUuid, in.Limit, in.Offset)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Internal Error: %v", err))
	}
	var result []*notification.Notification
	for _, ntf := range notifications {
		result = append(result, &notification.Notification{
			Id:     ntf.Id,
			Text:   ntf.Text,
			IsRead: ntf.IsRead,
		})
	}
	return &notification.NotificationOut{
		Notifications: result,
	}, nil
}

func (s *Service) MarkNotificationAsRead(ctx context.Context, in *notification.MarkNotificationAsReadIn) (*emptypb.Empty, error) {
	log.Println("MarkNotificationAsRead")
	userUuid := ctx.Value(config.KeyUUID).(string)
	err := s.dbR.MarkNotificationAsRead(ctx, userUuid, in.NotificationId)
	if err != nil {
		if err.Error() == "notification not found or already read" {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, fmt.Sprintf("Internal Error: %v", err))
	}
	return &emptypb.Empty{}, nil
}
