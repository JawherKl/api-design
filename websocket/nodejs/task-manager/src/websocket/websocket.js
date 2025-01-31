const WebSocket = require('ws');
const taskService = require('../services/taskService');

const setupWebSocket = (server) => {
  const wss = new WebSocket.Server({ server });

  wss.on('connection', (ws) => {
    console.log('New client connected');

    // Send initial tasks to the client
    taskService.getTasks().then((tasks) => {
      ws.send(JSON.stringify({ type: 'INITIAL_TASKS', data: tasks }));
    });

    // Handle incoming messages from the client
    ws.on('message', async (message) => {
      const { type, data } = JSON.parse(message);

      switch (type) {
        case 'CREATE_TASK':
          const newTask = await taskService.createTask(data);
          broadcast({ type: 'TASK_CREATED', data: newTask });
          break;

        case 'UPDATE_TASK':
          const updatedTask = await taskService.updateTask(data.id, data);
          broadcast({ type: 'TASK_UPDATED', data: updatedTask });
          break;

        case 'DELETE_TASK':
          await taskService.deleteTask(data.id);
          broadcast({ type: 'TASK_DELETED', data: { id: data.id } });
          break;
      }
    });

    // Broadcast messages to all clients
    const broadcast = (message) => {
      wss.clients.forEach((client) => {
        if (client.readyState === WebSocket.OPEN) {
          client.send(JSON.stringify(message));
        }
      });
    };

    ws.on('close', () => {
      console.log('Client disconnected');
    });
  });
};

module.exports = setupWebSocket;