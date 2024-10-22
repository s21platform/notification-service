package user

import (
	"context"
	"fmt"
	userproto "github.com/s21platform/user-proto/user-proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"notification-service/internal/config"
)

type Client struct {
	client userproto.UserServiceClient
}

func New(cfg *config.Config) *Client {
	connStr := fmt.Sprintf("%s:%s", cfg.User.Host, cfg.User.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	client := userproto.NewUserServiceClient(conn)
	return &Client{client: client}
}

func (c *Client) GetLoginByUuid(ctx context.Context, uuid string) (string, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uuid", uuid))
	resp, err := c.client.GetLoginByUUID(ctx, &userproto.GetLoginByUUIDIn{Uuid: uuid})
	if err != nil {
		return "", fmt.Errorf("failed to get login by uuid %s: %v", uuid, err)
	}
	return resp.Login, nil
}
