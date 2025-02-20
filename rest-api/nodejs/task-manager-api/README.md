# Task Management API

## 📌 Project Overview
The **Task Management API** is a backend service built with **Node.js, Express, and MongoDB**. It provides **user authentication (signup, login), task management (CRUD operations), and JWT-based authentication**.

## 🚀 Features
- **User Authentication** (Signup, Login, Token Generation)
- **JWT-based Authentication** for protected routes
- **Task Management** (Create, Read, Update, Delete tasks)
- **MongoDB Database** integration using Mongoose
- **Express.js Framework** for routing and middleware handling

## 🏗 Project Structure
```
task-manager-api/
│── src/
│   ├── config/           # Database & JWT config
│   ├── controllers/      # Business logic for tasks & users
│   ├── middlewares/      # Auth & validation middlewares
│   ├── models/           # Mongoose models (Task, User)
│   ├── routes/           # Route handlers
│   ├── utils/            # Helper functions (JWT, error handling)
│   ├── app.js            # Express application setup
│── server.js             # Main entry point
│── .env                  # Environment variables
│── package.json          # Dependencies and scripts
│── README.md             # Documentation
```

## 🔧 Installation & Setup
### **1. Clone the Repository**
```bash
git clone https://github.com/JawherKl/task-manager-api.git
cd task-manager-api
```

### **2. Install Dependencies**
```bash
npm install
```

### **3. Configure Environment Variables**
Create a `.env` file in the root directory and add:
```ini
PORT=5000
MONGO_URI=mongodb://localhost:27017/taskmanager
JWT_SECRET=your_jwt_secret_key
```

### **4. Start the Server**
```bash
npm run dev
```
The server will run at `http://localhost:5000`.

## 🔑 Authentication
The API uses **JWT-based authentication**. After login or signup, use the generated **Bearer token** to access protected routes.

### **Signup (POST /api/auth/signup)**
**Request:**
```json
{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "securepassword"
}
```
**Response:**
```json
{
  "user": { "_id": "12345", "username": "john_doe", "email": "john@example.com" },
  "token": "BearerTokenString"
}
```

### **Login (POST /api/auth/login)**
**Request:**
```json
{
  "email": "john@example.com",
  "password": "securepassword"
}
```
**Response:**
```json
{
  "user": { "_id": "12345", "username": "john_doe", "email": "john@example.com" },
  "token": "BearerTokenString"
}
```

## ✅ API Endpoints
| Method | Endpoint | Description | Auth Required |
|--------|---------|-------------|--------------|
| POST   | `/api/auth/signup` | Register a new user | No |
| POST   | `/api/auth/login`  | User login & get token | No |
| GET    | `/api/tasks` | Get all tasks | ✅ |
| POST   | `/api/tasks` | Create a new task | ✅ |
| PUT    | `/api/tasks/:id` | Update a task | ✅ |
| DELETE | `/api/tasks/:id` | Delete a task | ✅ |

**Authorization:** Include `Authorization: Bearer <token>` in the headers for protected routes.

## 📌 Next Steps
- ✅ Implement Refresh Token support
- ✅ Add Role-Based Access Control (RBAC)
- ✅ Deploy API to a cloud platform

## 📜 License
This project is licensed under the **MIT License**.

---

**🔥 Built with Node.js, Express, MongoDB & JWT.** Contributions are welcome! 🚀
