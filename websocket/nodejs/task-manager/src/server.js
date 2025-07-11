require('dotenv').config();
const express = require('express');
const http = require('http');
const { connectDB } = require('./config/db');
const setupWebSocket = require('./websocket/websocket');

const app = express();
const server = http.createServer(app);

// Connect to MongoDB
connectDB();

// Serve static files
app.use(express.static('public'));

// Set up WebSocket
setupWebSocket(server);

// Start the server
const PORT = process.env.PORT || 3000;
server.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});