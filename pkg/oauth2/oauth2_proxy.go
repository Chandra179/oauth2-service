package oauth2

import (
	"context"

	"golang.org/x/oauth2"
)

type Oauth2Proxy interface {
	Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
	Extra(key string, oauth2Token *oauth2.Token) interface{}
	S256ChallengeFromVerifier(verifier string) string
	AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
	TokenSource(ctx context.Context, t *oauth2.Token) oauth2.TokenSource
}

type Oauth2 struct {
	cfg *oauth2.Config
}

func NewOauth2(cfg *oauth2.Config) *Oauth2 {
	return &Oauth2{
		cfg: cfg,
	}
}

func (o *Oauth2) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return o.cfg.Exchange(ctx, code, opts...)
}

func (o *Oauth2) Extra(key string, oauth2Token *oauth2.Token) interface{} {
	return oauth2Token.Extra("id_token").(string)
}

func (o *Oauth2) S256ChallengeFromVerifier(verifier string) string {
	return oauth2.S256ChallengeFromVerifier(verifier)
}

func (o *Oauth2) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return o.cfg.AuthCodeURL(state, opts...)
}
func (o *Oauth2) TokenSource(ctx context.Context, t *oauth2.Token) oauth2.TokenSource {
	return o.cfg.TokenSource(ctx, t)
}
