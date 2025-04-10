package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepository_GetCountNotification(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := NewRepository(sqlxDB)

	ctx := context.Background()
	userUUID := "test-user-uuid"

	t.Run("success", func(t *testing.T) {
		expectedCount := int64(5)
		rows := sqlmock.NewRows([]string{"count"}).AddRow(expectedCount)
		mock.ExpectQuery(`SELECT COUNT\(id\) FROM push_notifications WHERE \(user_id = \$1 AND is_read = \$2\)`).
			WithArgs(userUUID, false).
			WillReturnRows(rows)

		count, err := repo.GetCountNotification(ctx, userUUID)
		require.NoError(t, err)
		assert.Equal(t, expectedCount, count)
	})
}

func TestRepository_GetNotifications(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := NewRepository(sqlxDB)

	ctx := context.Background()
	userUUID := "test-user-uuid"
	limit := int64(10)
	offset := int64(0)

	t.Run("success", func(t *testing.T) {
		now := time.Now()
		readAt := now.Add(-time.Hour)

		rows := sqlmock.NewRows([]string{"id", "notification", "is_read", "read_at", "created_at"}).
			AddRow(1, "Test notification 1", false, nil, now).
			AddRow(2, "Test notification 2", true, readAt, now.Add(-2*time.Hour))

		mock.ExpectQuery(`SELECT id, notification, is_read, read_at, created_at FROM push_notifications WHERE user_id = \$1 ORDER BY is_read ASC, created_at DESC LIMIT 10 OFFSET 0`).
			WithArgs(userUUID).
			WillReturnRows(rows)

		notifications, err := repo.GetNotifications(ctx, userUUID, limit, offset)
		require.NoError(t, err)
		assert.Len(t, notifications, 2)

		assert.Equal(t, int64(1), notifications[0].Id)
		assert.Equal(t, "Test notification 1", notifications[0].Text)
		assert.False(t, notifications[0].IsRead)
		assert.Nil(t, notifications[0].ReadAt)

		assert.Equal(t, int64(2), notifications[1].Id)
		assert.Equal(t, "Test notification 2", notifications[1].Text)
		assert.True(t, notifications[1].IsRead)
		assert.Equal(t, readAt.Unix(), notifications[1].ReadAt.Unix())
	})
}

func TestRepository_MarkNotificationAsRead(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := NewRepository(sqlxDB)

	ctx := context.Background()
	userUUID := "test-user-uuid"
	notificationID := int64(1)

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec(`UPDATE push_notifications SET is_read = \$1, read_at = NOW\(\) WHERE \(user_id = \$2 AND id = \$3 AND is_read = \$4\)`).
			WithArgs(true, userUUID, notificationID, false).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.MarkNotificationAsRead(ctx, userUUID, notificationID)
		require.NoError(t, err)
	})

	t.Run("notification not found", func(t *testing.T) {
		mock.ExpectExec(`UPDATE push_notifications SET is_read = \$1, read_at = NOW\(\) WHERE \(user_id = \$2 AND id = \$3 AND is_read = \$4\)`).
			WithArgs(true, userUUID, notificationID, false).
			WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.MarkNotificationAsRead(ctx, userUUID, notificationID)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "notification not found or already read")
	})
}
