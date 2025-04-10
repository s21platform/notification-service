package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/s21platform/notification-service/internal/config"
	"github.com/s21platform/notification-service/internal/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	connection *sqlx.DB
}

func (r *Repository) Close() {
	r.connection.Close()
}

func New(cfg *config.Config) *Repository {
	connectSourceStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Database, cfg.Postgres.Host, cfg.Postgres.Port)

	db, err := sqlx.Connect("postgres", connectSourceStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return &Repository{db}
}

func (r *Repository) GetCountNotification(ctx context.Context, userUuid string) (int64, error) {
	query, args, err := sq.Select(`COUNT(id)`).
		From("push_notifications").
		Where(sq.And{
			sq.Eq{"user_id": userUuid},
			sq.Eq{"is_read": false},
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("failed to build query: %v", err)
	}
	var count int
	err = r.connection.Get(&count, query, args...)
	if err != nil {
		return 0, fmt.Errorf("failed to execute query: %v", err)
	}
	return int64(count), nil
}

func (r *Repository) GetNotifications(ctx context.Context, userUuid string, limit int64, offset int64) ([]model.Notification, error) {
	query, args, err := sq.Select(`id`, `notification`, `is_read`, `read_at`, `created_at`).
		From(`push_notifications`).
		Where(sq.Eq{"user_id": userUuid}).
		OrderBy(`is_read ASC`, `created_at DESC`).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	var notifications []model.Notification
	err = r.connection.Select(&notifications, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	return notifications, nil
}

func (r *Repository) MarkNotificationAsRead(ctx context.Context, userUuid string, notificationId int64) error {
	query, args, err := sq.Update("push_notifications").
		Set("is_read", true).
		Set("read_at", sq.Expr("NOW()")).
		Where(sq.And{
			sq.Eq{"user_id": userUuid},
			sq.Eq{"id": notificationId},
			sq.Eq{"is_read": false},
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("failed to build query: %v", err)
	}

	result, err := r.connection.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("notification not found or already read")
	}

	return nil
}
