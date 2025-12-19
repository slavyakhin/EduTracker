package auth

import (
	"context"
	"log/slog"
	"time"

	"github.com/slavyakhin/EduTracker/services/sso/internal/domain/models"
)

type Auth struct {
	log          *slog.Logger
	userSaver    UserSaver
	userProvider UserProvider
	appProvider  AppProvider
	toketTTL     time.Duration
}

type UserSaver interface {
	SaveUser(
		ctx context.Context,
		email string,
		passHash []byte,
	) (uid int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, email string) (bool, error)
}

type AppProvider interface {
	App(ctx context.Context, appID int) (models.App, error)
}

// New returns a new instance of the Auth service.
func New(
	log *slog.Logger,
	userSaver UserSaver,
	userProvider UserProvider,
	appProvider AppProvider,
	tokenTTL time.Duration,
) *Auth {

	panic("Implement me!")
	// TODO: Implement

	return &Auth{
		log:          log,
		userSaver:    userSaver,
		userProvider: userProvider,
		appProvider:  appProvider,
		toketTTL:     tokenTTL,
	}
}

// Login check if user with given credentials exists in the system and returns access token.
//
// If user exists, but password is incorrect, returns error.
// If user doesn't exist, returns error.
func (a *Auth) Login(
	ctx context.Context,
	email string,
	password string,
	appID int,
) (string, error) {
	panic("Implement me!")
}

// RegisterNewUser registers new user in the system and returns user ID.
// If user with given username already exists, returns error.
func (s *Auth) RegisterNewUser(
	ctx context.Context,
	email string,
	pass string,
) (int64, string) {

	panic("Implement me!")
}

// IsAdmin checks if user is admin.
func (s *Auth) IsAdmin(
	ctx context.Context,
	userID int64,
) (bool, error) {

	panic("Implement me!")
}
