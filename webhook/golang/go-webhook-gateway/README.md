# 🚀 Go Webhook Gateway

A secure and extensible **Webhook Gateway** written in Go to receive and validate webhooks from external services like **GitHub**, **Stripe**, and **Twilio**, with MongoDB-based logging and modular processing support.

---

## 📌 Features

- ✅ Webhook receiver with dynamic source routes
- 🔐 Signature verification:
  - GitHub (HMAC-SHA256)
  - Stripe (HMAC-SHA256 with timestamp)
  - Twilio (HMAC-SHA1 with URL + form data)
- 🗃️ Event logging into MongoDB
- ⚙️ Extensible architecture for routing events to custom processors (Slack, Email, etc.)
- 🐳 Dockerized MongoDB for isolated development
- 🌱 Minimal dependencies with `Gin` and native Go libraries

---

## 🛠️ Tech Stack

| Layer     | Tech                          |
|-----------|-------------------------------|
| Language  | Go                            |
| Web       | [Gin](https://github.com/gin-gonic/gin)  |
| DB        | MongoDB (Docker container)    |
| Logging   | `log`                         |
| Security  | HMAC Signature Verifiers      |
| Tools     | `curl`, `ngrok`, `docker`     |

---

## 📁 Project Structure

```

go-webhook-gateway/
│
├── cmd/                  # Main entry point
│   └── main.go
├── config/               # (optional future: app config)
├── internal/
│   ├── handler/          # Webhook handler
│   ├── router/           # Route definitions
│   ├── storage/          # MongoDB connector
│   └── verifier/         # GitHub, Stripe, Twilio verifiers
├── models/               # Event struct model
├── test/                 # Unit tests (future)
├── .env                  # Config file (non-committed)
├── go.mod
└── README.md

````

---

## 🚀 Getting Started

### 1. Clone & Install Dependencies

```bash
git clone https://github.com/JawherKl/go-webhook-gateway.git
cd go-webhook-gateway
go mod tidy
````

---

### 2. Run MongoDB in Docker

```bash
docker run -d \
  --name webhook-mongo \
  -p 27017:27017 \
  -v webhook-mongo-data:/data/db \
  -e MONGO_INITDB_DATABASE=webhooks \
  mongo:7
```

---

### 3. Setup `.env` Configuration

```env
PORT=8080
MONGO_URI=mongodb://localhost:27017
GITHUB_WEBHOOK_SECRET=yourgithubsecret
STRIPE_WEBHOOK_SECRET=yourstripesecret
TWILIO_AUTH_TOKEN=yourtwiliotoken
```

---

### 4. Run the Gateway

```bash
go run cmd/main.go
```

Now the gateway listens on:
📬 `http://localhost:8080/webhook/:source`

---

## 🧪 Testing Webhooks

### ✅ GitHub Webhook Test (HMAC-SHA256)

```bash
payload='{"event":"push"}'
sig='sha256='$(echo -n $payload | openssl dgst -sha256 -hmac yourgithubsecret | sed 's/^.* //')

curl -X POST http://localhost:8080/webhook/github \
  -H "Content-Type: application/json" \
  -H "X-Hub-Signature-256: $sig" \
  -d "$payload"
```

---

### ✅ Stripe Webhook Test (HMAC-SHA256 + timestamp)

```bash
payload='{"id":"evt_test","type":"payment_intent.succeeded"}'
ts=$(date +%s)
sig='v1='$(echo -n "$ts.$payload" | openssl dgst -sha256 -hmac yourstripesecret | sed 's/^.* //')

curl -X POST http://localhost:8080/webhook/stripe \
  -H "Content-Type: application/json" \
  -H "Stripe-Signature: t=$ts,$sig" \
  -d "$payload"
```

---

### ✅ Twilio Webhook Test (HMAC-SHA1 + form-encoded)

Twilio signatures require the full request URL and form fields sorted. Best tested with **ngrok** + actual Twilio event.

---

## 📦 MongoDB Webhook Storage

Each verified event is stored in MongoDB:

```json
{
  "source": "stripe",
  "received_at": "2025-07-11T10:42:00Z",
  "headers": {
    "Stripe-Signature": "t=...,v1=..."
  },
  "payload": {
    "id": "evt_test",
    "type": "payment_intent.succeeded"
  }
}
```

---

## 🧩 Roadmap

* [x] GitHub webhook support
* [x] Stripe webhook support
* [x] Twilio webhook support
* [x] MongoDB event logging
* [ ] Add event processor system (Slack, Email, etc.)
* [ ] Add filters and search API
* [ ] Add Swagger documentation
* [ ] Add integration tests

---

## 📄 License

[MIT](LICENSE)

---

## 🙌 Acknowledgements

* [GitHub Webhook Docs](https://docs.github.com/en/webhooks)
* [Stripe Webhook Docs](https://stripe.com/docs/webhooks)
* [Twilio Webhook Docs](https://www.twilio.com/docs/usage/webhooks)

---

## 💡 Author

Made with ❤️ by [@JawherKl](https://github.com/JawherKl)
