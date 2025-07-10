package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"

    filepb "github.com/JawherKl/grpc-file-sharing/proto/file"
    metapb "github.com/JawherKl/grpc-file-sharing/proto/metadata"
    "github.com/JawherKl/grpc-file-sharing/services/file"
)

func main() {
    // Connect to metadata service
    conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Failed to connect to metadata service: %v", err)
    }
    defer conn.Close()

    metadataClient := metapb.NewMetadataServiceClient(conn)

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    filepb.RegisterFileServiceServer(grpcServer, &file.Server{
        MetadataClient: metadataClient,
    })

    log.Println("FileService running on :50051")
    grpcServer.Serve(lis)
}
