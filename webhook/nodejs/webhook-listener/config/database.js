const mongoose = require('mongoose');
require('dotenv').config(); // Load environment variables

const MONGO_URI = process.env.MONGO_URI || 'mongodb://localhost:27017/webhookDB';

async function connectDB() {
    try {
        await mongoose.connect(MONGO_URI, {
            useNewUrlParser: true,
            useUnifiedTopology: true,
        });
        console.log('✅ MongoDB Connected');
    } catch (error) {
        console.error('❌ MongoDB Connection Error:', error);
        process.exit(1); // Exit if connection fails
    }
}

module.exports = connectDB;
