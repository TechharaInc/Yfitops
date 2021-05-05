package service

import (
	"context"

	"golang.org/x/oauth2"
)

type ContextKey string

var (
	guildIDKey ContextKey = "guildIDKey"
	tokenKey   ContextKey = "tokenKey"
)

func SetGuildIDToContext(ctx context.Context, guildID string) context.Context {
	if guildID != "" {
		return context.WithValue(ctx, guildIDKey, guildID)
	}
	return ctx
}

func SetTokenToContext(ctx context.Context, token *oauth2.Token) context.Context {
	return context.WithValue(ctx, tokenKey, token)
}

func GetGuildIDFromContext(ctx context.Context) (string, bool) {
	v := ctx.Value(guildIDKey)
	gid, ok := v.(string)
	return gid, ok
}

func GetTokenFromContext(ctx context.Context) (*oauth2.Token, bool) {
	v := ctx.Value(tokenKey)
	tok, ok := v.(*oauth2.Token)
	return tok, ok
}
