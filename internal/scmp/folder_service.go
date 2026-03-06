package scmp

import (
	"context"

	pb "github.com/infodancer/scmp/proto/scmp/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// FolderServer implements the SCMP FolderService gRPC interface.
type FolderServer struct {
	pb.UnimplementedFolderServiceServer
}

func (s *FolderServer) ListFolders(_ context.Context, _ *pb.ListFoldersRequest) (*pb.ListFoldersResponse, error) {
	return nil, status.Error(codes.Unimplemented, "ListFolders not yet implemented")
}

func (s *FolderServer) CreateFolder(_ context.Context, req *pb.CreateFolderRequest) (*pb.CreateFolderResponse, error) {
	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}
	return nil, status.Error(codes.Unimplemented, "CreateFolder not yet implemented")
}

func (s *FolderServer) DeleteFolder(_ context.Context, req *pb.DeleteFolderRequest) (*pb.DeleteFolderResponse, error) {
	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}
	return nil, status.Error(codes.Unimplemented, "DeleteFolder not yet implemented")
}

func (s *FolderServer) RenameFolder(_ context.Context, req *pb.RenameFolderRequest) (*pb.RenameFolderResponse, error) {
	if req.GetOldName() == "" || req.GetNewName() == "" {
		return nil, status.Error(codes.InvalidArgument, "old_name and new_name are required")
	}
	return nil, status.Error(codes.Unimplemented, "RenameFolder not yet implemented")
}

func (s *FolderServer) MoveMessage(_ context.Context, req *pb.MoveMessageRequest) (*pb.MoveMessageResponse, error) {
	if len(req.GetMessageId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "message_id is required")
	}
	if req.GetDestinationFolder() == "" {
		return nil, status.Error(codes.InvalidArgument, "destination_folder is required")
	}
	return nil, status.Error(codes.Unimplemented, "MoveMessage not yet implemented")
}
