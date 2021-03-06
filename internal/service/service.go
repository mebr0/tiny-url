package service

import (
	"context"
	"github.com/mebr0/tiny-url/internal/cache"
	"github.com/mebr0/tiny-url/internal/domain"
	"github.com/mebr0/tiny-url/internal/repo"
	"github.com/mebr0/tiny-url/pkg/auth"
	"github.com/mebr0/tiny-url/pkg/hash"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Users interface {
	List(ctx context.Context) ([]domain.User, error)
}

type Auth interface {
	Register(ctx context.Context, toRegister domain.UserRegister) error
	Login(ctx context.Context, toLogin domain.UserLogin) (domain.Tokens, error)
}

type URLs interface {
	ListByOwner(ctx context.Context, owner primitive.ObjectID) ([]domain.URL, error)
	ListByOwnerAndExpiration(ctx context.Context, userId primitive.ObjectID, expired bool) ([]domain.URL, error)
	Create(ctx context.Context, toCreate domain.URLCreate) (domain.URL, error)
	Get(ctx context.Context, alias string) (domain.URL, error)
	GetByOwner(ctx context.Context, alias string, owner primitive.ObjectID) (domain.URL, error)
	Prolong(ctx context.Context, alias string, owner primitive.ObjectID, toProlong domain.URLProlong) (domain.URL, error)
	Delete(ctx context.Context, alias string, owner primitive.ObjectID) error
}

type Services struct {
	Users
	Auth
	URLs
}

func NewServices(repos *repo.Repos, caches *cache.Caches, hasher hash.PasswordHasher, tokenManager auth.TokenManager,
	urlEncoder hash.URLEncoder, accessTokenTTL time.Duration, aliasLength int, defaultExpiration int,
	urlCountLimit int) *Services {
	return &Services{
		Users: newUsersService(repos.Users),
		Auth:  newAuthService(repos.Users, hasher, tokenManager, accessTokenTTL),
		URLs:  newURLsService(repos.URLs, caches.URLs, urlEncoder, aliasLength, defaultExpiration, urlCountLimit),
	}
}
