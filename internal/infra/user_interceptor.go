package infra

import (
	"context"

	"github.com/s21platform/notification-service/internal/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	if info.FullMethod == "/NotificationService/SendVerificationCode" {
		return handler(ctx, req)
	}
	md, ok := metadata.FromIncomingContext(ctx)
	userIDs := md["uuid"]
	if !ok || len(userIDs) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "no uuid found in metadata")
	}
	ctx = context.WithValue(ctx, config.KeyUUID, userIDs[0])
	return handler(ctx, req)
}
