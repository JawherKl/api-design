const express = require('express');
const currencyRoutes = require('./routes/currencyRoutes');

const app = express();
app.use(express.json());

app.use('/api', currencyRoutes);

module.exports = app;
