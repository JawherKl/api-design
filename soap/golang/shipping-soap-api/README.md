# 📦 Shipping Order SOAP API (Go)

A simple powerful SOAP-based API written in Go for managing shipping orders. This project is ideal for learning how to build SOAP services using the `net/http` and `encoding/xml` standard libraries in Go.

---

## 🚀 Features

- ✅ Create new shipping orders with recipient details
- 🔍 Get shipping order status
- 🔄 Update delivery address (before delivery)
- ❌ Cancel shipping orders (if not yet delivered)
- 📦 In-memory storage for simplicity (extendable)
- 🧼 Clean SOAP envelope parsing using `encoding/xml`

---

## 🔧 Tech Stack

| Layer       | Technology         |
|-------------|--------------------|
| Language    | Go (Golang)        |
| Protocol    | SOAP/XML (v1.1)    |
| Web Server  | net/http           |
| XML Parser  | encoding/xml       |
| Storage     | In-Memory `map`    |
| UUIDs       | github.com/google/uuid |

---

## 📁 Project Structure

```
shipping-soap-api/
│
├── api/soap/               # SOAP HTTP handlers
│   └── order\_service.go
│
├── models/                 # Request/response structs
│   └── order.go
│   └── soap.go
│
├── storage/                # Order store (in-memory)
│   └── store.go
│
├── main.go                 # App entrypoint
├── go.mod / go.sum         # Go modules
└── README.md
````

---

## 📨 Sample SOAP Requests

### ▶️ Create Shipping Order

```xml
POST /soap
Content-Type: text/xml

<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <CreateShippingOrder>
      <RecipientName>Jane Smith</RecipientName>
      <DeliveryAddress>45 Queen Street, NY</DeliveryAddress>
      <ItemDescription>Smartphone</ItemDescription>
    </CreateShippingOrder>
  </soap:Body>
</soap:Envelope>
````

### ✅ Response

```xml
<CreateShippingOrderResponse>
  <TrackingNumber>abc123</TrackingNumber>
</CreateShippingOrderResponse>
```

---

### 🔍 Get Shipping Status

```xml
<GetShippingStatus>
  <TrackingNumber>abc123</TrackingNumber>
</GetShippingStatus>
```

---

### 🔄 Update Address

```xml
<UpdateDeliveryAddress>
  <TrackingNumber>abc123</TrackingNumber>
  <NewAddress>99 Park Avenue</NewAddress>
</UpdateDeliveryAddress>
```

---

### ❌ Cancel Order

```xml
<CancelShippingOrder>
  <TrackingNumber>abc123</TrackingNumber>
</CancelShippingOrder>
```

---

## 🧪 Running Locally

### 1. Clone & Install

```bash
git clone https://github.com/JawherKl/shipping-soap-api.git
cd shipping-soap-api
go mod tidy
```

### 2. Run Server

```bash
go run main.go
```

Server running at: [http://localhost:8080/soap](http://localhost:8080/soap)

---

## 🧠 Learnings

This project helps you understand:

* SOAP vs RESTful APIs
* XML parsing with Go's `encoding/xml`
* Building service-oriented Go APIs
* Clean architecture for legacy-compatible APIs

---

## 📜 To-Do / Improvements

* [ ] Add WSDL contract
* [ ] Add persistent storage (SQLite/PostgreSQL)
* [ ] Add authentication (SOAP headers)
* [ ] Add error codes and better fault messages
* [ ] Add tests with `httptest`

---

## 📄 License

MIT © [Jawher Kl](https://github.com/JawherKl)