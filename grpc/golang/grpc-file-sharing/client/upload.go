package main

import (
    "context"
    "log"
    "os"
    "time"
    "io"

    "google.golang.org/grpc"
    pb "github.com/JawherKl/grpc-file-sharing/proto/file"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Could not connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewFileServiceClient(conn)

    stream, err := client.Upload(context.Background())
    if err != nil {
        log.Fatalf("Could not upload: %v", err)
    }

    filePath := "testfile.txt"
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatalf("Could not open file: %v", err)
    }
    defer file.Close()

    buffer := make([]byte, 1024) // 1KB per chunk
    for {
        n, err := file.Read(buffer)
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatalf("Error reading file: %v", err)
        }

        err = stream.Send(&pb.FileChunk{
            Content:  buffer[:n],
            Filename: "testfile.txt",
            UserId:   "user-123",
        })
        if err != nil {
            log.Fatalf("Error sending chunk: %v", err)
        }

        time.Sleep(50 * time.Millisecond) // simulate delay
    }

    status, err := stream.CloseAndRecv()
    if err != nil {
        log.Fatalf("Upload failed: %v", err)
    }

    log.Printf("Upload status: %v - %s", status.Success, status.Message)
}
