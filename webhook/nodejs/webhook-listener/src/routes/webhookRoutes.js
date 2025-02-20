const express = require('express');
const router = express.Router();
const webhookController = require('../controllers/webhookController');
const verifySignature = require('../middlewares/verifySignature');

router.post('/receive', verifySignature, webhookController.handleWebhook);
router.get('/logs', webhookController.getWebhookEvents);

module.exports = router;
