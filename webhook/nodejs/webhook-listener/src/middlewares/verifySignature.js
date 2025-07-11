const crypto = require('crypto');
require('dotenv').config();

module.exports = (req, res, next) => {
    const receivedSignature = req.headers['x-webhook-signature'];
    const secret = process.env.WEBHOOK_SECRET || 'your-secret-key';

    if (!receivedSignature) {
        return res.status(403).json({ error: 'Signature missing' });
    }

    const payloadString = JSON.stringify(req.body, null, 0);

    const computedHash = crypto.createHmac('sha256', secret)
                               .update(payloadString)
                               .digest('hex');

    console.log("Received Signature: ", receivedSignature);
    console.log("Computed Signature: ", computedHash);
    console.log("Payload String: ", payloadString);

    if (computedHash !== receivedSignature) {
        return res.status(403).json({ error: 'Invalid signature' });
    }

    console.log("✅ Signature verified successfully!");
    next();
};
