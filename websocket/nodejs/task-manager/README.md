# Real-Time Collaborative Task Manager

A real-time task management application built with Node.js, Express, WebSocket, MongoDB, and Redis.

## Features
- Real-time collaboration: Multiple users can manage tasks simultaneously.
- Task management: Create, update, and delete tasks.
- Persistence: Tasks are stored in MongoDB.
- Caching: Frequently accessed tasks are cached using Redis.

## Tech Stack
- **Backend**: Node.js, Express, WebSocket
- **Database**: MongoDB
- **Caching**: Redis
- **Frontend**: HTML, JavaScript

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/JawherKl/api-design.git
   cd api-design
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Set up environment variables:
   Create a `.env` file and add the following:
   ```
   MONGO_URI=mongodb://ip-address:27017/taskmanager
   REDIS_URI=redis://ip-address:6379
   PORT=3000
   ```

4. Start the server:
   ```bash
   node src/server.js
   ```

5. Open your browser and navigate to `http://localhost:3000`.

## API Endpoints
- **WebSocket**: `ws://localhost:3000` (for real-time updates)

## Contributing
Pull requests are welcome! For major changes, please open an issue first to discuss the changes.

## License
This project is licensed under the MIT License.
```

---
<!--
### **Next Steps**
1. Add authentication (e.g., JWT) to secure the application.
2. Use a frontend framework like React or Vue.js for a better user experience.
3. Deploy the application using platforms like Heroku, Vercel, or AWS.
-->