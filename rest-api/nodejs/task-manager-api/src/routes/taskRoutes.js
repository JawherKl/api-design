const express = require("express");
const { createTask, getAllTasks, getTaskById, updateTask, deleteTask } = require("../controllers/taskControllers");
const validateTask = require("../middlewares/validateTask");
const authMiddleware = require("../middlewares/authMiddleware");

const router = express.Router();

router.post("/", authMiddleware, validateTask, createTask);
router.get("/", authMiddleware, getAllTasks);
router.get("/:id", authMiddleware, getTaskById);
router.put("/:id", authMiddleware, validateTask, updateTask);
router.delete("/:id", authMiddleware, deleteTask);

module.exports = router;
