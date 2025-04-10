package service

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/s21platform/notification-service/internal/config"
	"github.com/s21platform/notification-service/internal/model"
	"github.com/s21platform/notification-service/pkg/notification"
)

func TestService_GetNotificationCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockDbRepo(ctrl)
	service := New(mockRepo)
	ctx := context.WithValue(context.Background(), config.KeyUUID, "test-user-uuid")

	t.Run("success", func(t *testing.T) {
		expectedCount := int64(5)
		mockRepo.EXPECT().GetCountNotification(ctx, "test-user-uuid").Return(expectedCount, nil)

		result, err := service.GetNotificationCount(ctx, &emptypb.Empty{})

		assert.NoError(t, err)
		assert.Equal(t, expectedCount, result.Count)
	})

	t.Run("repository error", func(t *testing.T) {
		mockRepo.EXPECT().GetCountNotification(ctx, "test-user-uuid").Return(int64(0), assert.AnError)

		result, err := service.GetNotificationCount(ctx, &emptypb.Empty{})

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestService_GetNotification(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockDbRepo(ctrl)
	service := New(mockRepo)
	ctx := context.WithValue(context.Background(), config.KeyUUID, "test-user-uuid")

	now := time.Now()
	mockNotifications := []model.Notification{
		{
			Id:        1,
			Text:      "Test notification 1",
			IsRead:    false,
			CreatedAt: now,
		},
		{
			Id:        2,
			Text:      "Test notification 2",
			IsRead:    true,
			ReadAt:    &now,
			CreatedAt: now.Add(-time.Hour),
		},
	}

	t.Run("success", func(t *testing.T) {
		input := &notification.NotificationIn{Limit: 10, Offset: 0}
		mockRepo.EXPECT().GetNotifications(ctx, "test-user-uuid", input.Limit, input.Offset).
			Return(mockNotifications, nil)

		result, err := service.GetNotification(ctx, input)

		assert.NoError(t, err)
		assert.Len(t, result.Notifications, 2)

		// Проверяем первое уведомление
		assert.Equal(t, mockNotifications[0].Id, result.Notifications[0].Id)
		assert.Equal(t, mockNotifications[0].Text, result.Notifications[0].Text)
		assert.Equal(t, mockNotifications[0].IsRead, result.Notifications[0].IsRead)

		// Проверяем второе уведомление
		assert.Equal(t, mockNotifications[1].Id, result.Notifications[1].Id)
		assert.Equal(t, mockNotifications[1].Text, result.Notifications[1].Text)
		assert.Equal(t, mockNotifications[1].IsRead, result.Notifications[1].IsRead)
	})

	t.Run("empty result", func(t *testing.T) {
		input := &notification.NotificationIn{Limit: 10, Offset: 100}
		mockRepo.EXPECT().GetNotifications(ctx, "test-user-uuid", input.Limit, input.Offset).
			Return([]model.Notification{}, nil)

		result, err := service.GetNotification(ctx, input)

		assert.NoError(t, err)
		assert.Empty(t, result.Notifications)
	})

	t.Run("repository error", func(t *testing.T) {
		input := &notification.NotificationIn{Limit: 10, Offset: 0}
		mockRepo.EXPECT().GetNotifications(ctx, "test-user-uuid", input.Limit, input.Offset).
			Return(nil, assert.AnError)

		result, err := service.GetNotification(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestService_MarkNotificationAsRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockDbRepo(ctrl)
	service := New(mockRepo)
	ctx := context.WithValue(context.Background(), config.KeyUUID, "test-user-uuid")

	t.Run("success", func(t *testing.T) {
		input := &notification.MarkNotificationAsReadIn{NotificationId: 1}
		mockRepo.EXPECT().MarkNotificationAsRead(ctx, "test-user-uuid", input.NotificationId).
			Return(nil)

		_, err := service.MarkNotificationAsRead(ctx, input)
		assert.NoError(t, err)
	})

	t.Run("notification not found", func(t *testing.T) {
		input := &notification.MarkNotificationAsReadIn{NotificationId: 999}
		mockRepo.EXPECT().MarkNotificationAsRead(ctx, "test-user-uuid", input.NotificationId).
			Return(ErrNotificationNotFound)

		_, err := service.MarkNotificationAsRead(ctx, input)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "notification not found")
	})

	t.Run("repository error", func(t *testing.T) {
		input := &notification.MarkNotificationAsReadIn{NotificationId: 1}
		mockRepo.EXPECT().MarkNotificationAsRead(ctx, "test-user-uuid", input.NotificationId).
			Return(assert.AnError)

		_, err := service.MarkNotificationAsRead(ctx, input)
		assert.Error(t, err)
	})
}
