// Code generated by ogen, DO NOT EDIT.

package fanbox

import (
	"context"
	"net/http"

	"github.com/go-faster/errors"
)

// SecuritySource is provider of security values (tokens, passwords, etc.).
type SecuritySource interface {
	// CsrfToken provides csrfToken security value.
	// You must set CSTF Token in header.
	CsrfToken(ctx context.Context, operationName string) (CsrfToken, error)
	// SessionId provides sessionId security value.
	// You must set Session ID in cookie.
	SessionId(ctx context.Context, operationName string) (SessionId, error)
}

func (s *Client) securityCsrfToken(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.CsrfToken(ctx, operationName)
	if err != nil {
		return errors.Wrap(err, "security source \"CsrfToken\"")
	}
	req.Header.Set("X-CSRF-Token", t.APIKey)
	return nil
}
func (s *Client) securitySessionId(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.SessionId(ctx, operationName)
	if err != nil {
		return errors.Wrap(err, "security source \"SessionId\"")
	}
	req.AddCookie(&http.Cookie{
		Name:  "FANBOXSESSID",
		Value: t.APIKey,
	})
	return nil
}
