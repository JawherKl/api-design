package file

import (
    "context"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "time"
    "log"

    filepb "github.com/JawherKl/grpc-file-sharing/proto/file"
    metapb "github.com/JawherKl/grpc-file-sharing/proto/metadata"

    "google.golang.org/grpc"
)

type Server struct {
    filepb.UnimplementedFileServiceServer
    MetadataClient metapb.MetadataServiceClient
}

func (s *Server) Upload(stream filepb.FileService_UploadServer) error {
    var (
        file     *os.File
        filePath string
        totalSize int64
        filename string
        userId string
    )

    for {
        chunk, err := stream.Recv()
        if err == io.EOF {
            // Done receiving file → close and save metadata
            if file != nil {
                file.Close()

                // Save metadata
                metadata := &metapb.FileMetadata{
                    UserId:     userId,
                    Filename:   filename,
                    Size:       totalSize,
                    UploadedAt: time.Now().Format(time.RFC3339),
                }

                _, err := s.MetadataClient.SaveMetadata(context.Background(), metadata)
                if err != nil {
                    return stream.SendAndClose(&filepb.UploadStatus{
                        Success: false,
                        Message: fmt.Sprintf("File saved, but failed to store metadata: %v", err),
                    })
                }

                return stream.SendAndClose(&filepb.UploadStatus{
                    Success: true,
                    Message: "File uploaded and metadata saved successfully",
                })
            }

            return stream.SendAndClose(&filepb.UploadStatus{
                Success: false,
                Message: "No file received",
            })
        }
        if err != nil {
            return err
        }

        // First chunk
        if file == nil {
            filename = chunk.Filename
            userId = chunk.UserId
            uploadDir := filepath.Join("uploads", userId)
            os.MkdirAll(uploadDir, os.ModePerm)
            filePath = filepath.Join(uploadDir, filename)
            file, err = os.Create(filePath)
            if err != nil {
                return stream.SendAndClose(&filepb.UploadStatus{
                    Success: false,
                    Message: "Failed to create file: " + err.Error(),
                })
            }
        }

        // Write chunk
        n, err := file.Write(chunk.Content)
        if err != nil {
            return stream.SendAndClose(&filepb.UploadStatus{
                Success: false,
                Message: "Failed to write chunk: " + err.Error(),
            })
        }
        totalSize += int64(n)
    }
}
