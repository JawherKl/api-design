const WebhookEvent = require('../models/WebhookEvent');

exports.handleWebhook = async (req, res) => {
    try {
        console.log("Received Webhook:", req.body);
        
        // Save webhook event to database
        const event = new WebhookEvent({ payload: req.body });
        await event.save();
        
        res.status(200).json({ message: "Webhook received successfully" });
    } catch (error) {
        console.error("Error handling webhook:", error);
        res.status(500).json({ error: "Internal Server Error" });
    }
};

exports.getWebhookEvents = async (req, res) => {
    try {
        const events = await WebhookEvent.find().sort({ receivedAt: -1 });
        res.json(events);
    } catch (error) {
        res.status(500).json({ error: "Internal Server Error" });
    }
};
