package scmp

import (
	"context"

	pb "github.com/infodancer/scmp/proto/scmp/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthServer implements the SCMP AuthService gRPC interface.
type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

func (s *AuthServer) Authenticate(_ context.Context, req *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	if req.GetAddress() == "" {
		return nil, status.Error(codes.InvalidArgument, "address is required")
	}
	if req.GetCredential() == nil {
		return nil, status.Error(codes.InvalidArgument, "credential is required")
	}
	return nil, status.Error(codes.Unimplemented, "Authenticate not yet implemented")
}

func (s *AuthServer) Logout(_ context.Context, _ *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Logout not yet implemented")
}
