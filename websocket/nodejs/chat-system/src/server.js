const express = require('express');
const http = require('http');
const WebSocket = require('ws');
const path = require('path');

const app = express();
const server = http.createServer(app);
const wss = new WebSocket.Server({ server });

// Serve static files from the "public" folder
app.use(express.static(path.join(__dirname, '../public')));

// WebSocket connection handler
wss.on('connection', (ws) => {
  console.log('New client connected');

  // Send a welcome message to the client
  ws.send('Welcome to the WebSocket server!');

  // Handle incoming messages from the client
  ws.on('message', (message) => {
    console.log(`Received: ${message}`);

    // Broadcast the message to all connected clients
    wss.clients.forEach((client) => {
      if (client !== ws && client.readyState === WebSocket.OPEN) {
        client.send(`User said: ${message}`);
      }
    });
  });

  // Handle client disconnection
  ws.on('close', () => {
    console.log('Client disconnected');
  });
});

// Start the server
const PORT = process.env.PORT || 3000;
server.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});