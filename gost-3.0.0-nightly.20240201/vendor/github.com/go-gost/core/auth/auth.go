package auth

import "context"

type Options struct{}

type Option func(opts *Options)

// Authenticator is an interface for user authentication.
// 认证器：通用接口
type Authenticator interface {
	Authenticate(ctx context.Context, user, password string, opts ...Option) (id string, ok bool)
}

type authenticatorGroup struct {
	authers []Authenticator
}

func AuthenticatorGroup(authers ...Authenticator) Authenticator {
	return &authenticatorGroup{
		authers: authers,
	}
}

func (p *authenticatorGroup) Authenticate(ctx context.Context, user, password string, opts ...Option) (string, bool) {
	if len(p.authers) == 0 {
		return "", false
	}
	for _, auther := range p.authers {
		if auther == nil {
			continue
		}

		if id, ok := auther.Authenticate(ctx, user, password, opts...); ok {
			return id, ok
		}
	}
	return "", false
}
