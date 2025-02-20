const axios = require("axios");
const crypto = require("crypto");
require("dotenv").config();

const WEBHOOK_URL = "https://9e8b-193-95-3-82.ngrok-free.app/webhooks/receive"; // Use ngrok URL in real tests
const SECRET_KEY = process.env.WEBHOOK_SECRET || "my-secret-key";

const webhookPayload = {
    event: "user.signup",
    user: "john_doe",
    email: "john@example.com",
};

// Compute the signature using HMAC SHA256
const signature = crypto
    .createHmac("sha256", SECRET_KEY)
    .update(JSON.stringify(webhookPayload))
    .digest("hex");

// Send webhook with computed signature
async function sendWebhook() {
    try {
        const response = await axios.post(WEBHOOK_URL, webhookPayload, {
            headers: {
                "Content-Type": "application/json",
                "x-webhook-signature": signature,
            },
        });

        console.log("✅ Webhook sent successfully:", response.data);
    } catch (error) {
        console.error("❌ Webhook failed:", error.response ? error.response.data : error.message);
    }
}

// Run webhook test
sendWebhook();
