package scmp

import (
	"context"

	pb "github.com/infodancer/scmp/proto/scmp/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PolicyServer implements the SCMP PolicyService gRPC interface.
type PolicyServer struct {
	pb.UnimplementedPolicyServiceServer
}

func (s *PolicyServer) SetPullPreference(_ context.Context, req *pb.SetPullPreferenceRequest) (*pb.SetPullPreferenceResponse, error) {
	if req.GetSenderDomain() == "" {
		return nil, status.Error(codes.InvalidArgument, "sender_domain is required")
	}
	return nil, status.Error(codes.Unimplemented, "SetPullPreference not yet implemented")
}

func (s *PolicyServer) GetPullPreference(_ context.Context, req *pb.GetPullPreferenceRequest) (*pb.GetPullPreferenceResponse, error) {
	if req.GetSenderDomain() == "" {
		return nil, status.Error(codes.InvalidArgument, "sender_domain is required")
	}
	return nil, status.Error(codes.Unimplemented, "GetPullPreference not yet implemented")
}

func (s *PolicyServer) ListPullPreferences(_ context.Context, _ *pb.ListPullPreferencesRequest) (*pb.ListPullPreferencesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "ListPullPreferences not yet implemented")
}

func (s *PolicyServer) SetDefaultPolicy(_ context.Context, _ *pb.SetDefaultPolicyRequest) (*pb.SetDefaultPolicyResponse, error) {
	return nil, status.Error(codes.Unimplemented, "SetDefaultPolicy not yet implemented")
}

func (s *PolicyServer) GetDefaultPolicy(_ context.Context, _ *pb.GetDefaultPolicyRequest) (*pb.GetDefaultPolicyResponse, error) {
	return nil, status.Error(codes.Unimplemented, "GetDefaultPolicy not yet implemented")
}
