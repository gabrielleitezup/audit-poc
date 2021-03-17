package util

import (
	"context"
	"net/http"
)

type ContextKey string

const (
	AuthContextKey      = ContextKey("jwt")
	EntityContextKey    = ContextKey("table")
	UserAgentContextKey = ContextKey("user-agent")
	UserIpContextKey    = ContextKey("user-ip")
)

func FillContext(r *http.Request, user, entity string) context.Context {
	ctxUser := context.WithValue(r.Context(), AuthContextKey, user)
	ctxUserAgent := context.WithValue(ctxUser, UserAgentContextKey, r.UserAgent())
	ctxRemoteAddress := context.WithValue(ctxUserAgent, UserIpContextKey, r.RemoteAddr)
	ctxTableAddress := context.WithValue(ctxRemoteAddress, EntityContextKey, entity)
	return ctxTableAddress
}
