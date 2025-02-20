const soap = require("soap");
const express = require("express");
const axios = require("axios");
const fs = require("fs");
const mongoose = require("mongoose");
require("dotenv").config();
const connectDB = require("../config/database");
const Transaction = require("../models/Transaction");

const EXCHANGE_API_URL = "https://api.exchangerate-api.com/v4/latest/USD";

let exchangeRates = {};

// 1️⃣ Connect to MongoDB
connectDB();

// 2️⃣ Fetch exchange rates
const updateExchangeRates = async () => {
    try {
        console.log("Fetching latest exchange rates...");
        const response = await axios.get(EXCHANGE_API_URL);
        exchangeRates = response.data.rates;
        console.log("Exchange rates updated successfully.");
    } catch (error) {
        console.error("Error fetching exchange rates:", error.message);
    }
};

updateExchangeRates();
setInterval(updateExchangeRates, 60 * 60 * 1000); // Refresh every hour

// 3️⃣ Define SOAP service
const service = {
    CurrencyExchangeService: {
        CurrencyExchangePort: {
            ConvertCurrency: async function (args) {
                const { from, to, amount } = args;

                if (!exchangeRates[from] || !exchangeRates[to]) {
                    throw new Error(`Invalid currency: ${from} or ${to}`);
                }

                const usdAmount = parseFloat(amount) / exchangeRates[from];
                const convertedAmount = usdAmount * exchangeRates[to];

                console.log(`Converted ${amount} ${from} to ${convertedAmount.toFixed(2)} ${to}`);

                // 4️⃣ Save transaction to MongoDB
                try {
                    const transaction = new Transaction({
                        from,
                        to,
                        amount,
                        convertedAmount,
                        rate: exchangeRates[to],
                    });
                    await transaction.save();
                    console.log("Transaction saved to MongoDB.");
                } catch (err) {
                    console.error("Error saving transaction:", err.message);
                }

                return { result: convertedAmount.toFixed(2) };
            },
        },
    },
};

// 5️⃣ Load WSDL and start server
const xml = fs.readFileSync("./wsdl/currency.wsdl", "utf8");

const app = express();
app.use("/wsdl", (req, res) => res.send(xml));

app.listen(3001, () => {
    soap.listen(app, "/soap", service, xml);
    console.log("SOAP server running on http://localhost:3001/soap");
});
