const Task = require('../models/taskModel');
const redisClient = require('../config/db').redisClient;

// Get all tasks (with caching)
const getTasks = async () => {
  const cachedTasks = await redisClient.get('tasks');
  if (cachedTasks) {
    return JSON.parse(cachedTasks);
  }

  const tasks = await Task.find();
  await redisClient.set('tasks', JSON.stringify(tasks), { EX: 60 }); // Cache for 60 seconds
  return tasks;
};

// Create a new task
const createTask = async (taskData) => {
  const newTask = await Task.create(taskData);
  await redisClient.del('tasks'); // Invalidate cache
  return newTask;
};

// Update a task
const updateTask = async (id, taskData) => {
  const updatedTask = await Task.findByIdAndUpdate(id, taskData, { new: true });
  await redisClient.del('tasks'); // Invalidate cache
  return updatedTask;
};

// Delete a task
const deleteTask = async (id) => {
  await Task.findByIdAndDelete(id);
  await redisClient.del('tasks'); // Invalidate cache
};

module.exports = { getTasks, createTask, updateTask, deleteTask };