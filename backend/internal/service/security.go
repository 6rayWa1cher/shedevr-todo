package service

import (
	"context"
	"github.com/6rayWa1cher/shedevr-todo/backend/pkg/oas"
)

type SecurityService interface {
	oas.SecurityHandler
	GetCurrentRemoteUser(ctx context.Context) (string, bool)
}

type securityService struct {
	SecurityService
}

func NewSecurityService() SecurityService {
	return securityService{}
}

const ctxRemoteUserKey = "remoteUser"

func (s securityService) HandleRemoteUserAuth(
	ctx context.Context, _ string, t oas.RemoteUserAuth,
) (context.Context, error) {
	return context.WithValue(ctx, ctxRemoteUserKey, t.APIKey), nil
}

func (s securityService) GetCurrentRemoteUser(ctx context.Context) (string, bool) {
	out, ok := ctx.Value(ctxRemoteUserKey).(string)
	return out, ok
}
