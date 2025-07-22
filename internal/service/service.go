package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/s21platform/notification-service/pkg/notification"

	"github.com/s21platform/notification-service/internal/config"
)

var ErrNotificationNotFound = errors.New("notification not found or already read")

type Service struct {
	notification.UnimplementedNotificationServiceServer
	dbR           DbRepo
	emailS        EmailSender
	verificationS VerificationCodeSender
	vecS          VerificationEduCodeSender
}

func New(repo DbRepo, emailSender EmailSender, verificationSender VerificationCodeSender, vecS VerificationEduCodeSender) *Service {
	return &Service{
		dbR:           repo,
		emailS:        emailSender,
		verificationS: verificationSender,
		vecS:          vecS,
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

func (s *Service) MarkNotificationsAsRead(ctx context.Context, in *notification.MarkNotificationsAsReadIn) (*emptypb.Empty, error) {
	log.Println("MarkNotificationAsRead")
	userUuid := ctx.Value(config.KeyUUID).(string)
	err := s.dbR.MarkNotificationsAsRead(ctx, userUuid, in.NotificationIds)
	if err != nil {
		if err.Error() == "notification not found or already read" {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, fmt.Sprintf("Internal Error: %v", err))
	}
	return &emptypb.Empty{}, nil
}

func (s *Service) SendVerificationCode(ctx context.Context, in *notification.SendVerificationCodeIn) (*emptypb.Empty, error) {
	log.Println("SendVerificationCode")

	email := in.GetEmail()
	code := in.GetCode()

	log.Println("email", email)
	log.Println("code", code)

	if email == "" || code == "" {
		return nil, status.Error(codes.InvalidArgument, "email and code are required")
	}

	// Используем специализированный сервис для отправки верификационного кода
	if err := s.verificationS.SendVerificationCode(email, code); err != nil {
		log.Printf("failed to send verification code: %v", err)
		return nil, status.Error(codes.Internal, fmt.Sprintf("Internal Error: %v", err))
	}

	log.Printf("verification code sent to %s", email)
	return &emptypb.Empty{}, nil
}

func (s *Service) SendEduCode(ctx context.Context, in *notification.SendEduCodeIn) (*emptypb.Empty, error) {
	log.Println("SendEduCode")

	if in.Email == "" || in.Code == "" {
		return nil, status.Error(codes.InvalidArgument, "email and code are required")
	}

	if !strings.Contains(in.Email, "@student.21-school.ru") {
		return nil, status.Error(codes.InvalidArgument, "invalid email")
	}

	err := s.vecS.SendVerificationCode(in.Email, in.Code)
	if err != nil {
		log.Printf("failed to send verification edu code: %v", err)
		return nil, status.Error(codes.Internal, fmt.Sprintf("Internal Error: %v", err))
	}
	log.Printf("verification edu code sent to %s", in.Email)
	return &emptypb.Empty{}, nil
}
