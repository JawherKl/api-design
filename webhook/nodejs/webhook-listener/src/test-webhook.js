const axios = require('axios');
const crypto = require('crypto');

const secret = 'your-secret-key';
const payload = { message: "Automated Webhook Test" };

const payloadString = JSON.stringify(payload);  

const signature = crypto.createHmac('sha256', secret)
                        .update(payloadString)
                        .digest('hex');

console.log("Generated Signature:", signature);

axios.post('http://localhost:3000/webhooks/receive', payload, {
    headers: {
        'Content-Type': 'application/json',
        'x-webhook-signature': signature
    }
})
.then(response => console.log("Success:", response.data))
.catch(error => console.error("Error:", error.response?.data || error.message));
