package contextHandler

import (
	"context"
	"time"
)

// sessionKey      = new(cfg.GetSecuritySessionKey_Session())
// userKeyKey      = new(cfg.GetSecuritySessionKey_UserKey())
// userCodeKey     = new(cfg.GetSecuritySessionKey_UserCode())
// tokenKey        = new(cfg.GetSecuritySessionKey_Token())
// expiryPeriodKey = new(cfg.GetSecuritySessionKey_ExpiryPeriod())

func GetUserCode(ctx context.Context) string {
	return ctx.Value(userCodeKey).(string)
}

func GetUserKey(ctx context.Context) string {
	return ctx.Value(userKeyKey).(string)
}

func GetSessionID(ctx context.Context) string {
	return ctx.Value(sessionIDKey).(string)
}

func GetSessionToken(ctx context.Context) any {
	return ctx.Value(tokenKey)
}

func GetSessionExpiry(ctx context.Context) time.Time {
	return ctx.Value(expiryPeriodKey).(time.Time)
}

// Setters

func SetSessionID(ctx context.Context, sessionID string) context.Context {
	return context.WithValue(ctx, sessionIDKey, sessionID)
}

func SetUserKey(ctx context.Context, userKey string) context.Context {
	return context.WithValue(ctx, userKeyKey, userKey)
}

func SetUserCode(ctx context.Context, userCode string) context.Context {
	return context.WithValue(ctx, userCodeKey, userCode)
}

func SetSessionToken(ctx context.Context, token any) context.Context {
	return context.WithValue(ctx, tokenKey, token)
}

func SetSessionExpiry(ctx context.Context, expiry time.Time) context.Context {
	return context.WithValue(ctx, expiryPeriodKey, expiry)
}
