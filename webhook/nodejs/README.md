# Webhook Listener with Node.js & MongoDB

## 📌 Project Overview
This project is a **webhook listener** built with **Node.js, Express, and MongoDB**. It securely receives webhook events, verifies signatures, and stores them in MongoDB.

## ✨ Features
- ✅ **Webhook Listener**: Listens for webhook events via a POST API.
- 🔐 **Signature Verification**: Ensures webhook authenticity.
- 💾 **MongoDB Integration**: Stores webhook events in a database.
- 🚀 **Docker Support**: Easily run MongoDB in a container.
- 📜 **Logging & Error Handling**: Provides useful logs and error management.

---

## ⚡ Installation & Setup
### 1️⃣ Clone the Repository
```sh
git clone https://github.com/YOUR_USERNAME/webhook-listener.git
cd webhook-listener
```

### 2️⃣ Install Dependencies
```sh
npm install
```

### 3️⃣ Configure Environment Variables
Create a `.env` file in the root directory:
```
PORT=3000
MONGO_URI=mongodb://localhost:27017/webhookDB
WEBHOOK_SECRET=your_secret_key
```

### 4️⃣ Start MongoDB
If MongoDB is **not installed**, use Docker:
```sh
docker run -d --name mongodb -p 27017:27017 mongo
```

If MongoDB is **installed locally**, start it:
```sh
sudo systemctl start mongod
```

### 5️⃣ Run the Server
Created account at ngrok and follow the guide to setup ngrok webhook server
```sh
ngrok http 3000

node index.js
```

---

## 📌 API Usage
### 🎯 Webhook Receiver
**Endpoint:** `POST /webhooks/receive`

**Headers:**
```
Content-Type: application/json
x-webhook-signature: <generated-signature>
```

**Payload Example:**
```json
{
  "event": "user.signup",
  "user": "john_doe",
  "email": "john@example.com"
}
```

**Success Response:**
```json
{
  "success": true,
  "message": "Webhook stored"
}
```

---

## 🛠 Project Structure
```
webhook-listener/
│── models/
│   ├── WebhookEvent.js  # Mongoose schema for storing webhooks
│── routes/
│   ├── webhook.js       # Webhook route handler
│── .env                 # Environment variables
│── database.js          # MongoDB connection setup
│── index.js             # Main Express server file
│── package.json         # Project dependencies
```

---

## 🔍 Troubleshooting
### 1️⃣ MongoDB Connection Issues
Check if MongoDB is running:
```sh
systemctl status mongod
```
Or start a MongoDB container:
```sh
docker run -d --name mongodb -p 27017:27017 mongo
```

### 2️⃣ Webhook Not Storing Data
Check the logs for errors:
```sh
node index.js
```
Ensure your `.env` file has the correct `MONGO_URI`.

### 3️⃣ Signature Verification Failing
Ensure the `WEBHOOK_SECRET` is correctly set and used in signature generation.

---

## 📜 License
This project is licensed under the MIT License.
