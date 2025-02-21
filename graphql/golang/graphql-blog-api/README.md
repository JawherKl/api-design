# GraphQL Blog API (Golang + MongoDB)

## 📌 Overview
The **GraphQL Blog API** is a backend service built with **Golang** and **MongoDB**, using **GraphQL** to manage blog posts and comments. This project demonstrates how to:

- Set up a **GraphQL server** in Golang using `gqlgen`.
- Connect to **MongoDB** for data storage.
- Implement **GraphQL queries and mutations**.
- Structure a Golang project for a **GraphQL API**.

---

## 📁 Project Structure
```
graphql-blog-api/
│── config/           # Database connection setup
│── graph/            # GraphQL schema & resolvers
│── models/           # MongoDB models
│── routes/           # Router setup
│── server/           # Server initialization
│── main.go           # Entry point
│── go.mod            # Dependencies
│── go.sum            # Checksums
│── .env              # Environment variables
```

---

## ⚡ Getting Started

### 1️⃣ Clone the Repository
```sh
git clone https://github.com/yourusername/graphql-blog-api.git
cd graphql-blog-api
```

### 2️⃣ Install Dependencies
```sh
go mod tidy
```

### 3️⃣ Set Up Environment Variables
Create a `.env` file and add your MongoDB connection string:
```env
MONGODB_URI=mongodb://localhost:27017
```

### 4️⃣ Run the Application
```sh
go run main.go
```
The server will start on **http://localhost:8080/**.

---

## 🗄️ Database Model
### **Post Model**
```go
type Post struct {
  ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
  Title    string             `bson:"title" json:"title"`
  Content  string             `bson:"content" json:"content"`
  Author   string             `bson:"author" json:"author"`
  Comments []Comment          `bson:"comments,omitempty" json:"comments,omitempty"`
}
```

### **Comment Model**
```go
type Comment struct {
  ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
  PostID  primitive.ObjectID `bson:"postId" json:"postId"`
  Author  string             `bson:"author" json:"author"`
  Content string             `bson:"content" json:"content"`
}
```

---

## 📜 GraphQL Schema

```graphql
type Post {
  id: ID!
  title: String!
  content: String!
  author: String!
  comments: [Comment]
}

type Comment {
  id: ID!
  postId: ID!
  author: String!
  content: String!
}

type Query {
  posts: [Post]
  post(id: ID!): Post
}

type Mutation {
  createPost(title: String!, content: String!, author: String!): Post
  deletePost(id: ID!): String
}
```

---

## 🚀 API Usage
### **1️⃣ Create a Blog Post**
#### **Mutation**
```graphql
mutation {
  createPost(title: "First Post", content: "This is my first blog post!", author: "John Doe") {
    id
    title
    content
    author
  }
}
```

### **2️⃣ Fetch All Blog Posts**
#### **Query**
```graphql
query {
  posts {
    id
    title
    content
    author
  }
}
```

---

## 📌 Next Steps
✅ Add **comment mutations** (`createComment`, `deleteComment`)<br>
✅ Implement **error handling** and **input validation**<br>
✅ Add **JWT authentication** for secure user access<br>
✅ Deploy to a **cloud server** or **Dockerize** the project<br>

---

## 💡 Contributing
Feel free to submit issues or pull requests to improve the project!

---

## 📜 License
This project is licensed under the MIT License.
