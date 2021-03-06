package repo

import (
	"context"
	"github.com/mebr0/tiny-url/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

//go:generate mockgen -source=repo.go -destination=mocks/mock.go

type Users interface {
	List(ctx context.Context) ([]domain.User, error)
	Create(ctx context.Context, user domain.User) (primitive.ObjectID, error)
	GetByCredentials(ctx context.Context, email, password string) (domain.User, error)
	UpdateLastLogin(ctx context.Context, id primitive.ObjectID, lastLogin time.Time) error
}

type URLs interface {
	ListByOwner(ctx context.Context, userId primitive.ObjectID) ([]domain.URL, error)
	ListByOwnerAndExpiration(ctx context.Context, userId primitive.ObjectID, expired bool) ([]domain.URL, error)
	Create(ctx context.Context, url domain.URL) (string, error)
	Get(ctx context.Context, alias string) (domain.URL, error)
	GetByOriginalAndOwner(ctx context.Context, original string, owner primitive.ObjectID) (domain.URL, error)
	Prolong(ctx context.Context, alias string, toProlong domain.URLProlong) error
	Delete(ctx context.Context, alias string) error
}

type Repos struct {
	Users Users
	URLs  URLs
}

func NewRepos(db *mongo.Database) *Repos {
	return &Repos{
		Users: newUsersRepo(db),
		URLs:  newURLsRepo(db),
	}
}
