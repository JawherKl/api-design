const Joi = require("joi");

const taskSchema = Joi.object({
    title: Joi.string().required(),
    description: Joi.string().optional(),
    status: Joi.string().valid("pending", "in-progress", "completed").default("pending"),
    dueDate: Joi.date().optional(),
});

const validateTask = (req, res, next) => {
    const { error } = taskSchema.validate(req.body);
    if (error) return res.status(400).json({ message: error.details[0].message });
    next();
};

module.exports = validateTask;
