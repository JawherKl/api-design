const mongoose = require('mongoose');

const WebhookEventSchema = new mongoose.Schema({
    payload: { type: Object, required: true },
    receivedAt: { type: Date, default: Date.now }
});

module.exports = mongoose.model('WebhookEvent', WebhookEventSchema);
