package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/slavyakhin/EduTracker/services/sso/internal/app/grpc"
	"github.com/slavyakhin/EduTracker/services/sso/internal/services/auth"
	"github.com/slavyakhin/EduTracker/services/sso/internal/storage/postgres"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storageURI string,
	tokenTTL time.Duration,
) *App {
	storage, err := postgres.New(storageURI)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
