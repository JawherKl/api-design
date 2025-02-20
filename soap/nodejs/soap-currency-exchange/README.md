# Currency Exchange SOAP API

## 📌 Overview
This project is a **SOAP-based Currency Exchange API** built with **Node.js, Express, and MongoDB**. It provides real-time currency conversion by fetching exchange rates from an external API and logs transactions in a MongoDB database.

## 🚀 Features
- **SOAP API** for currency conversion
- **Real-time exchange rates** fetched from `https://api.exchangerate-api.com`
- **MongoDB integration** to store conversion transactions
- **Express.js & SOAP server** for API handling
- **Automatic rate updates** every hour

## 🛠️ Installation & Setup

### 1️⃣ Clone the Repository
```sh
git clone https://github.com/JawherKl/api-design.git
cd soap/nodejs/soap-currency-exchange
```

### 2️⃣ Install Dependencies
```sh
npm install
```

### 3️⃣ Configure Environment Variables
Create a `.env` file in the root directory:
```
MONGO_URI=mongodb://localhost:27017/currency-exchange
```

Replace with your MongoDB connection string if using MongoDB Atlas.

### 4️⃣ Start the Server
```sh
node src/services/soapServer.js
```

The SOAP server will run on:
```
http://localhost:3001/soap
```

## 📝 API Usage

### 1️⃣ SOAP Request Format
Send a **SOAP request** to the API endpoint using **Postman** or **cURL**:

```xml
<?xml version="1.0"?>
<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
  <Body>
    <ConvertCurrency xmlns="http://localhost:3001/soap">
      <from>USD</from>
      <to>EUR</to>
      <amount>100</amount>
    </ConvertCurrency>
  </Body>
</Envelope>
```

### 2️⃣ Test with cURL
```sh
curl -X POST http://localhost:3001/soap \  
-H "Content-Type: text/xml" \  
-d '<?xml version="1.0"?>
<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
  <Body>
    <ConvertCurrency xmlns="http://localhost:3001/soap">
      <from>USD</from>
      <to>EUR</to>
      <amount>100</amount>
    </ConvertCurrency>
  </Body>
</Envelope>'
```

### 3️⃣ Expected Response
```xml
<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
  <Body>
    <ConvertCurrencyResponse>
      <result>91.00</result>
    </ConvertCurrencyResponse>
  </Body>
</Envelope>
```

## 🗂 Database Structure
MongoDB stores each transaction in the `transactions` collection with this structure:

```json
{
  "_id": "65d5c5e4abc123456",
  "from": "USD",
  "to": "EUR",
  "amount": 100,
  "convertedAmount": 91,
  "rate": 0.91,
  "timestamp": "2025-02-20T14:30:00Z"
}
```

### Fetch Saved Transactions (MongoDB)
```sh
mongo
use currency-exchange
db.transactions.find().pretty()
```

## 📂 Project Structure
```
📦 currency-exchange-soap-api
├── 📁 src
│   ├── 📁 config
│   │   ├── database.js       # MongoDB connection setup
│   ├── 📁 models
│   │   ├── Transaction.js    # Mongoose schema for transactions
│   ├── 📁 services
│   │   ├── soapServer.js     # SOAP API implementation
├── 📁 wsdl
│   ├── currency.wsdl         # WSDL file defining the SOAP service
├── .env                      # Environment variables
├── package.json              # Dependencies & scripts
├── README.md                 # Documentation
```

## 📜 License
This project is licensed under the **MIT License**.

## 👨‍💻 Author
**JawherKl** – [GitHub](https://github.com/JawherKl)

---
🚀 **Enjoy using the Currency Exchange SOAP API!**

