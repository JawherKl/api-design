package metadata

import (
    "context"
    "time"

    "gorm.io/gorm"

    pb "github.com/JawherKl/grpc-file-sharing/proto/metadata"
)

type Server struct {
    pb.UnimplementedMetadataServiceServer
    DB *gorm.DB
}

type File struct {
    ID         string `gorm:"primaryKey"`
    UserID     string
    Filename   string
    Size       int64
    UploadedAt time.Time
}

func (s *Server) SaveMetadata(ctx context.Context, in *pb.FileMetadata) (*pb.SaveStatus, error) {
    f := File{
        ID:         in.UserId + "_" + in.Filename,
        UserID:     in.UserId,
        Filename:   in.Filename,
        Size:       in.Size,
        UploadedAt: time.Now(),
    }

    if err := s.DB.Create(&f).Error; err != nil {
        return &pb.SaveStatus{Success: false, Message: err.Error()}, nil
    }

    return &pb.SaveStatus{Success: true, Message: "Metadata saved"}, nil
}

func (s *Server) GetMetadata(ctx context.Context, in *pb.FileRequest) (*pb.FileMetadata, error) {
    var f File
    err := s.DB.First(&f, "user_id = ? AND filename = ?", in.UserId, in.Filename).Error
    if err != nil {
        return nil, err
    }

    return &pb.FileMetadata{
        UserId:     f.UserID,
        Filename:   f.Filename,
        Size:       f.Size,
        UploadedAt: f.UploadedAt.Format(time.RFC3339),
    }, nil
}

func (s *Server) ListFiles(ctx context.Context, in *pb.UserRequest) (*pb.FileList, error) {
    var files []File
    err := s.DB.Find(&files, "user_id = ?", in.UserId).Error
    if err != nil {
        return nil, err
    }

    var result []*pb.FileMetadata
    for _, f := range files {
        result = append(result, &pb.FileMetadata{
            UserId:     f.UserID,
            Filename:   f.Filename,
            Size:       f.Size,
            UploadedAt: f.UploadedAt.Format(time.RFC3339),
        })
    }

    return &pb.FileList{Files: result}, nil
}
