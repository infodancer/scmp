package scmp

import (
	"context"

	pb "github.com/infodancer/scmp/proto/scmp/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// KeyServer implements the SCMP KeyService gRPC interface.
type KeyServer struct {
	pb.UnimplementedKeyServiceServer
}

func (s *KeyServer) SubmitPublicKey(_ context.Context, req *pb.SubmitPublicKeyRequest) (*pb.SubmitPublicKeyResponse, error) {
	if req.GetAlgorithm() == "" {
		return nil, status.Error(codes.InvalidArgument, "algorithm is required")
	}
	if len(req.GetPublicKey()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "public_key is required")
	}
	return nil, status.Error(codes.Unimplemented, "SubmitPublicKey not yet implemented")
}

func (s *KeyServer) GetSignedKey(_ context.Context, _ *pb.GetSignedKeyRequest) (*pb.GetSignedKeyResponse, error) {
	return nil, status.Error(codes.Unimplemented, "GetSignedKey not yet implemented")
}

func (s *KeyServer) RotateKey(_ context.Context, req *pb.RotateKeyRequest) (*pb.RotateKeyResponse, error) {
	if req.GetAlgorithm() == "" {
		return nil, status.Error(codes.InvalidArgument, "algorithm is required")
	}
	if len(req.GetNewPublicKey()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "new_public_key is required")
	}
	return nil, status.Error(codes.Unimplemented, "RotateKey not yet implemented")
}

func (s *KeyServer) RevokeKey(_ context.Context, _ *pb.RevokeKeyRequest) (*pb.RevokeKeyResponse, error) {
	return nil, status.Error(codes.Unimplemented, "RevokeKey not yet implemented")
}

func (s *KeyServer) LookupKey(_ context.Context, req *pb.LookupKeyRequest) (*pb.LookupKeyResponse, error) {
	if req.GetAddress() == "" {
		return nil, status.Error(codes.InvalidArgument, "address is required")
	}
	return nil, status.Error(codes.Unimplemented, "LookupKey not yet implemented")
}

func (s *KeyServer) MonitorFingerprint(_ context.Context, _ *pb.MonitorFingerprintRequest) (*pb.MonitorFingerprintResponse, error) {
	return nil, status.Error(codes.Unimplemented, "MonitorFingerprint not yet implemented")
}
