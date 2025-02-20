const express = require('express');
const axios = require('axios');

const router = express.Router();

router.get('/convert', async (req, res) => {
    const { from, to } = req.query;

    if (!from || !to) {
        return res.status(400).json({ error: "Missing 'from' or 'to' parameter" });
    }

    try {
        // Fetch exchange rates from a public API (fix this if needed)
        const response = await axios.get(`https://api.exchangerate-api.com/v4/latest/${from}`);
        
        if (!response.data.rates[to]) {
            return res.status(400).json({ error: "Invalid currency code" });
        }

        res.json({ rate: response.data.rates[to] });
    } catch (error) {
        console.error("Exchange rate fetch error:", error.message);
        res.status(500).json({ error: "Error fetching exchange rate" });
    }
});

module.exports = router;
