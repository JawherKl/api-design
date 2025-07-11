<h1 align="center">📡 gRPC File Sharing API (Go)</h1>

<p align="center">
  ⚡️ Upload and store files with metadata using gRPC microservices in Go.
</p>

---

## 📦 Project Structure

```

grpc-file-sharing/
├── client/                # gRPC clients (e.g. upload)
├── cmd/                   # Main entry points for services
│   ├── file-service/
│   └── metadata-service/
├── proto/                 # Protobuf files + generated code
├── services/              # Service implementations
├── uploads/               # Uploaded files (saved by user)
├── go.mod
└── README.md

````

---

## 🚀 Features

- 📤 **File Upload** via streaming gRPC
- 🧠 **Metadata Storage** in PostgreSQL (filename, user, size, upload time)
- 📂 **Auto-directory creation** by user ID
- 🔗 **Inter-service communication** between FileService & MetadataService
- 💡 Simple gRPC client (`upload.go`) for testing

---

## 🛠️ Tech Stack

| Layer             | Tech/Tool                    |
|------------------|------------------------------|
| Language         | Go 1.24+                      |
| Protocol         | gRPC + Protocol Buffers       |
| Storage          | FileSystem (`uploads/`)      |
| Metadata DB      | PostgreSQL + GORM             |
| Dependencies     | `grpc-go`, `protoc-gen-go`, `gorm` |
| DevOps           | (Optional) Docker / Compose   |

---

## 🧩 Services Overview

### 📁 FileService (`:50051`)
Handles file upload as a gRPC stream, stores content, then calls `MetadataService`.

- Endpoint: `Upload(stream FileChunk) returns UploadStatus`
- Saves to: `/uploads/{userId}/{filename}`

### 🧠 MetadataService (`:50052`)
Stores metadata (user, filename, size, upload date) into PostgreSQL.

- Methods:
  - `SaveMetadata()`
  - `GetMetadata()`
  - `ListFiles()`

---

## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/JawherKl/api-design.git
cd grpc/golang/grpc-file-sharing
````

### 2. Setup PostgreSQL

```bash
docker run --name pg-metadata -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres
# Then create database manually:
# CREATE DATABASE metadata_db;
```

### 3. Generate gRPC code

```bash
protoc --go_out=. --go-grpc_out=. proto/file.proto
protoc --go_out=. --go-grpc_out=. proto/metadata.proto
```

### 4. Run the services

```bash
# Terminal 1
go run cmd/metadata-service/main.go

# Terminal 2
go run cmd/file-service/main.go
```

### 5. Run the client

```bash
echo "This is a test file" > testfile.txt
go run client/upload.go
```

---

## ✅ Result

* File saved to: `uploads/user-123/testfile.txt`
* Metadata saved in PostgreSQL:

```sql
SELECT * FROM files;
```

---

## 📁 Proto Example

```proto
// file.proto
service FileService {
  rpc Upload (stream FileChunk) returns (UploadStatus);
}

message FileChunk {
  bytes content = 1;
  string filename = 2;
  string userId = 3;
}

message UploadStatus {
  bool success = 1;
  string message = 2;
}
```

---

## 📦 TODOs

* [ ] Add `ListFiles` client
* [ ] Add Docker & Docker Compose
* [ ] Add authentication
* [ ] Add unit + integration tests
* [ ] Deploy to cloud / gRPC Gateway REST

---

## 🤝 Contribution

Pull requests welcome. Please open an issue first to discuss what you would like to change.

---

## 📜 License

This project is open-source and available under the [MIT License](LICENSE).

---

## 💬 Contact

Made with 💚 by [Jawher Kallel](https://github.com/JawherKl)
