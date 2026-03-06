package scmp

import (
	"io"
	"log"

	pb "github.com/infodancer/scmp/proto/scmp/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NotificationServer implements the SCMP NotificationService gRPC interface.
// Provides bidirectional streaming for real-time push from server to client.
type NotificationServer struct {
	pb.UnimplementedNotificationServiceServer
}

func (s *NotificationServer) Subscribe(stream pb.NotificationService_SubscribeServer) error {
	// Read client events in a goroutine; push server events on the same stream.
	// For v0.0.1, we just log client events and return Unimplemented
	// until the event dispatch system is built.

	for {
		clientEvent, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return status.Errorf(codes.Internal, "recv: %v", err)
		}

		switch evt := clientEvent.GetEvent().(type) {
		case *pb.ClientEvent_Ack:
			log.Printf("scmp: client ack for sequence %d", evt.Ack.GetSequence())
		case *pb.ClientEvent_RegisterHashes:
			log.Printf("scmp: client registered %d hashes", len(evt.RegisterHashes.GetHashes()))
		}
	}
}
