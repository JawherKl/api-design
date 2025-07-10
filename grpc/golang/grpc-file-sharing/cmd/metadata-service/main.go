package main

import (
    "log"
    "net"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "google.golang.org/grpc"

    pb "github.com/JawherKl/grpc-file-sharing/proto/metadata"
    "github.com/JawherKl/grpc-file-sharing/services/metadata"
)

func main() {
    dsn := "host=localhost user=postgres password=secret dbname=metadata_db port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to DB: %v", err)
    }

    db.AutoMigrate(&metadata.File{})

    lis, err := net.Listen("tcp", ":50052")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterMetadataServiceServer(grpcServer, &metadata.Server{DB: db})

    log.Println("MetadataService gRPC server started on :50052")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
