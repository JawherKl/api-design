const soap = require('soap');

const url = 'http://currencyconverter.kowabunga.net/converter.asmx?WSDL';

async function getExchangeRate(from, to) {
    const client = await soap.createClientAsync(url);
    const result = await client.GetConversionAmountAsync({ CurrencyFrom: from, CurrencyTo: to, RateDate: new Date().toISOString() });
    return result[0];
}

module.exports = { getExchangeRate };
