require('dotenv').config();
const express = require('express');
const bodyParser = require('body-parser');
const cors = require('cors');
const helmet = require('helmet');
const webhookRoutes = require('../src/routes/webhookRoutes');
const connectDB = require('../config/database');

const app = express();
app.use(bodyParser.json());
app.use(helmet());
app.use(cors());
connectDB();

app.use('/webhooks', webhookRoutes);

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`Webhook listener running on port ${PORT}`);
});
