package scmp

import (
	"context"

	pb "github.com/infodancer/scmp/proto/scmp/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MessageServer implements the SCMP MessageService gRPC interface.
type MessageServer struct {
	pb.UnimplementedMessageServiceServer
}

func (s *MessageServer) SubmitMessage(_ context.Context, req *pb.SubmitMessageRequest) (*pb.SubmitMessageResponse, error) {
	if req.GetDestination() == "" {
		return nil, status.Error(codes.InvalidArgument, "destination is required")
	}
	if len(req.GetEnvelope()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "envelope is required")
	}
	if len(req.GetEncryptedPayload()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "encrypted_payload is required")
	}
	return nil, status.Error(codes.Unimplemented, "SubmitMessage not yet implemented")
}

func (s *MessageServer) ListMessages(_ context.Context, _ *pb.ListMessagesRequest) (*pb.ListMessagesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "ListMessages not yet implemented")
}

func (s *MessageServer) FetchMessage(_ context.Context, req *pb.FetchMessageRequest) (*pb.FetchMessageResponse, error) {
	if len(req.GetMessageId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "message_id is required")
	}
	return nil, status.Error(codes.Unimplemented, "FetchMessage not yet implemented")
}

func (s *MessageServer) AcceptMessage(_ context.Context, req *pb.AcceptMessageRequest) (*pb.AcceptMessageResponse, error) {
	if len(req.GetMessageId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "message_id is required")
	}
	return nil, status.Error(codes.Unimplemented, "AcceptMessage not yet implemented")
}

func (s *MessageServer) DeleteMessage(_ context.Context, req *pb.DeleteMessageRequest) (*pb.DeleteMessageResponse, error) {
	if len(req.GetMessageId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "message_id is required")
	}
	return nil, status.Error(codes.Unimplemented, "DeleteMessage not yet implemented")
}

func (s *MessageServer) ListPending(_ context.Context, _ *pb.ListPendingRequest) (*pb.ListPendingResponse, error) {
	return nil, status.Error(codes.Unimplemented, "ListPending not yet implemented")
}
