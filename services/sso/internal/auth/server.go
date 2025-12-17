package auth

import (
	"context"
	"net/mail"

	ssov1 "github.com/slavyakhin/EduTracker/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{})
}

const (
	emptyValue = 0
)

func (s *serverAPI) Register(
	ctx context.Context,
	req *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	panic("Implement me")
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	// validation
	if _, err := mail.ParseAddress(req.GetEmail()); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid email")
	}
	if len(req.GetPassword()) < 8 {
		return nil, status.Error(codes.InvalidArgument, "password length is less than 8")
	}
	if req.GetAppId() == emptyValue {
		return nil, status.Error(codes.InvalidArgument, "app_id is required")
	}

	// TODO: implement login via auth service

	return &ssov1.LoginResponse{
		Token: req.GetEmail(),
	}, nil
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	req *ssov1.IsAdminRequest,
) (*ssov1.IsAdminResponse, error) {
	panic("Implement me")
}
