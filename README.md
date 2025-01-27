# API Design Study Repository 🚀

![api-design](https://github.com/JawherKl/api-design/blob/main/docs/images/api-design.gif)

Welcome to the **API Design Study Repository**! This project is a hands-on exploration of API design patterns, covering **REST API**, **GraphQL**, **gRPC**, **SOAP**, **Webhook** and **Websocket** implementations across four popular frameworks: **Node.js**, **Golang**, **Symfony**, and **Spring Boot**. 

---

## 📂 Project Structure

```
api-design-study/
├── README.md
├── graphql/
│   ├── nodejs/
│   │   ├── package.json
│   │   ├── src/
│   │   │   ├── resolvers/
│   │   │   └── schema/
│   │   └── app.js
│   ├── golang/
│   │   ├── go.mod
│   │   ├── main.go
│   │   ├── resolvers/
│   │   └── schema/
│   ├── symfony/
│   │   ├── composer.json
│   │   └── src/
│   │       ├── GraphQL/
│   │       ├── Entity/
│   │       └── config/
│   └── spring-boot/
│       ├── pom.xml
│       └── src/
│           ├── main/
│           │   ├── java/
│           │   │   ├── resolvers/
│           │   │   └── schema/
├── grpc/
│   ├── nodejs/
│   │   ├── package.json
│   │   ├── src/
│   │   │   ├── proto/
│   │   │   ├── server.js
│   │   │   └── client.js
│   ├── golang/
│   │   ├── go.mod
│   │   ├── proto/
│   │   ├── server/
│   │   └── client/
│   ├── symfony/
│   │   ├── composer.json
│   │   └── src/
│   │       ├── GRPC/
│   │       ├── Entity/
│   │       └── config/
│   └── spring-boot/
│       ├── pom.xml
│       └── src/
│           ├── main/
│           │   ├── java/
│           │   │   ├── grpc/
│           │   │   ├── server/
│           │   │   └── client/
├── rest-api/
│   ├── nodejs/
│   │   ├── package.json
│   │   ├── src/
│   │   │   ├── controllers/
│   │   │   ├── routes/
│   │   │   └── app.js
│   ├── golang/
│   │   ├── go.mod
│   │   ├── main.go
│   │   ├── handlers/
│   │   └── routes/
│   ├── symfony/
│   │   ├── composer.json
│   │   └── src/
│   │       ├── Controller/
│   │       ├── Entity/
│   │       └── config/
│   └── spring-boot/
│       ├── pom.xml
│       └── src/
│           ├── main/
│           │   ├── java/
│           │   │   ├── controllers/
│           │   │   ├── services/
│           │   │   └── Application.java
├── soap/
│   ├── nodejs/
│   │   ├── package.json
│   │   ├── src/
│   │   │   ├── controllers/
│   │   │   ├── routes/
│   │   │   └── app.js
│   ├── golang/
│   │   ├── go.mod
│   │   ├── main.go
│   │   ├── handlers/
│   │   └── routes/
│   ├── symfony/
│   │   ├── composer.json
│   │   └── src/
│   │       ├── Controller/
│   │       ├── Entity/
│   │       └── config/
│   └── spring-boot/
│       ├── pom.xml
│       └── src/
│           ├── main/
│           │   ├── java/
│           │   │   ├── controllers/
│           │   │   ├── services/
│           │   │   └── Application.java
├── webhook/
│   ├── nodejs/
│   │   ├── package.json
│   │   ├── src/
│   │   │   ├── controllers/
│   │   │   ├── routes/
│   │   │   └── app.js
│   ├── golang/
│   │   ├── go.mod
│   │   ├── main.go
│   │   ├── handlers/
│   │   └── routes/
│   ├── symfony/
│   │   ├── composer.json
│   │   └── src/
│   │       ├── Controller/
│   │       ├── Entity/
│   │       └── config/
│   └── spring-boot/
│       ├── pom.xml
│       └── src/
│           ├── main/
│           │   ├── java/
│           │   │   ├── controllers/
│           │   │   ├── services/
│           │   │   └── Application.java
├── websocket/
│   ├── nodejs/
│   │   ├── package.json
│   │   ├── src/
│   │   │   ├── controllers/
│   │   │   ├── routes/
│   │   │   └── app.js
│   ├── golang/
│   │   ├── go.mod
│   │   ├── main.go
│   │   ├── handlers/
│   │   └── routes/
│   ├── symfony/
│   │   ├── composer.json
│   │   └── src/
│   │       ├── Controller/
│   │       ├── Entity/
│   │       └── config/
│   └── spring-boot/
│       ├── pom.xml
│       └── src/
│           ├── main/
│           │   ├── java/
│           │   │   ├── controllers/
│           │   │   ├── services/
│           │   │   └── Application.java
└── docs/
    ├── REST-API.md
    ├── GraphQL.md
    ├── gRPC.md
    ├── soap.md
    ├── webhook.md
    ├── websocket.md
    └── framework-comparisons.md
```

---

## ✨ Features

- **REST API**: Comprehensive study of HTTP methods, request/response handling, and best practices.
- **GraphQL**: Efficient data retrieval with GraphQL schemas, resolvers, and queries.
- **gRPC**: Implementation of Protobuf-based gRPC server-client communication.
- **Webhook**: Automated HTTP callbacks for real-time event notifications.
- **WebSocket**: Full-duplex communication channels for real-time data exchange over a single TCP connection.

---

## 🌐 Frameworks Covered

- **Node.js**: JavaScript runtime for building scalable APIs.
- **Golang**: Lightweight and fast language for robust API services.
- **Symfony**: PHP framework with structured components and rich tools.
- **Spring Boot**: Java-based framework tailored for microservices and enterprise APIs.

---

## 🚀 Getting Started

### Prerequisites

Ensure you have the following installed:
- **Node.js**
- **Go**
- **PHP**
- **Java** (with Maven)
- **Docker** (for optional containerized services)

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/JawherKl/api-design-study.git
   ```

2. Navigate to a specific API type and framework:
   ```bash
   cd api-design-study/rest-api/nodejs
   ```

3. Install dependencies and start the server:
   ```bash
   npm install
   npm start
   ```

---

## 📚 Documentation

- [REST API Design](docs/REST-API.md)
- [GraphQL API Design](docs/GraphQL.md)
- [gRPC API Design](docs/gRPC.md)
- [SOAP API Design](docs/SOAP.md)
- [webhook API Design](docs/Webhook.md)
- [websocket API Design](docs/Websocket.md)
- [Framework Comparisons](docs/framework-comparisons.md)

---

## 🤝 Contribution

We welcome contributions! Feel free to:
- Fork this repository
- Raise issues
- Submit pull requests to enhance the repository

---

### 🌟 Happy Learning and Building! 🌟

## Stargazers over time 🌟

[![Stargazers over time](https://starchart.cc/JawherKl/api-design.svg?variant=adaptive)](https://starchart.cc/JawherKl/api-design)
