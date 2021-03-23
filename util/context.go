package util

import (
	"context"
	"net/http"
)

type ContextKey string

const (
	AuthContextKey      = ContextKey("jwt")
	UserAgentContextKey = ContextKey("user-agent")
	UserIpContextKey    = ContextKey("user-ip")
)

func FillContext(r *http.Request, user string) context.Context {
	ctxUser := context.WithValue(r.Context(), AuthContextKey, user)
	ctxUserAgent := context.WithValue(ctxUser, UserAgentContextKey, r.UserAgent())
	ctxRemoteAddress := context.WithValue(ctxUserAgent, UserIpContextKey, r.RemoteAddr)
	return ctxRemoteAddress
}
