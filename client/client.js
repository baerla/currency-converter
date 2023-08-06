const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const path = require('path');

// Load the protobuf file dynamically from path
const protoPath = path.join(__dirname, "..", "pb_schemas", "currency.proto");
const protoDefinition = protoLoader.loadSync(protoPath,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });
const currencyProto = grpc.loadPackageDefinition(protoDefinition).currency;

// Create gRPC client
const client = new currencyProto.CurrencyConverter('localhost:50051', grpc.credentials.createInsecure());

// Function to convert currency
function convertCurrency(fromCurrency, toCurrency, amount) {
    const request = {
        from_currency: fromCurrency,
        to_currency: toCurrency,
        amount: amount
    };

    console.log(request)

    client.Convert(request, (error, response) => {
        if (!error) {
            console.log(`Converted ${amount} ${fromCurrency} to ${response.converted_amount} ${toCurrency}`);
        } else {
            console.error("Error: ", error.details);
        }
    });

}

// Convert currency
convertCurrency('USD', 'EUR', 100);