// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/ogenerrors"
)

// SecurityHandler is handler for security parameters.
type SecurityHandler interface {
	// HandleRemoteUserAuth handles RemoteUserAuth security.
	HandleRemoteUserAuth(ctx context.Context, operationName string, t RemoteUserAuth) (context.Context, error)
}

func findAuthorization(h http.Header, prefix string) (string, bool) {
	v, ok := h["Authorization"]
	if !ok {
		return "", false
	}
	for _, vv := range v {
		scheme, value, ok := strings.Cut(vv, " ")
		if !ok || !strings.EqualFold(scheme, prefix) {
			continue
		}
		return value, true
	}
	return "", false
}

func (s *Server) securityRemoteUserAuth(ctx context.Context, operationName string, req *http.Request) (context.Context, bool, error) {
	var t RemoteUserAuth
	const parameterName = "X-Remote-User"
	value := req.Header.Get(parameterName)
	if value == "" {
		return ctx, false, nil
	}
	t.APIKey = value
	rctx, err := s.sec.HandleRemoteUserAuth(ctx, operationName, t)
	if errors.Is(err, ogenerrors.ErrSkipServerSecurity) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}
	return rctx, true, err
}

// SecuritySource is provider of security values (tokens, passwords, etc.).
type SecuritySource interface {
	// RemoteUserAuth provides RemoteUserAuth security value.
	RemoteUserAuth(ctx context.Context, operationName string) (RemoteUserAuth, error)
}

func (s *Client) securityRemoteUserAuth(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.RemoteUserAuth(ctx, operationName)
	if err != nil {
		return errors.Wrap(err, "security source \"RemoteUserAuth\"")
	}
	req.Header.Set("X-Remote-User", t.APIKey)
	return nil
}