package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/infodancer/scmp/internal/scmp"
	pb "github.com/infodancer/scmp/proto/scmp/v1"
	"google.golang.org/grpc"
)

func main() {
	grpcAddr := envOrDefault("SCMP_GRPC_ADDR", ":9500")

	grpcListener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", grpcAddr, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &scmp.AuthServer{})
	pb.RegisterMessageServiceServer(grpcServer, &scmp.MessageServer{})
	pb.RegisterFolderServiceServer(grpcServer, &scmp.FolderServer{})
	pb.RegisterKeyServiceServer(grpcServer, &scmp.KeyServer{})
	pb.RegisterPolicyServiceServer(grpcServer, &scmp.PolicyServer{})
	pb.RegisterNotificationServiceServer(grpcServer, &scmp.NotificationServer{})

	go func() {
		log.Printf("scmp: gRPC server listening on %s", grpcAddr)
		if err := grpcServer.Serve(grpcListener); err != nil {
			log.Fatalf("grpc serve: %v", err)
		}
	}()

	fmt.Println("scmp: Secure Client Messaging Protocol server running")
	fmt.Printf("  gRPC: %s\n", grpcAddr)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	log.Println("scmp: shutting down")
	grpcServer.GracefulStop()
}

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
