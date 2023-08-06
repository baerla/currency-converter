# Currency Converter

A little currency converter to try gRPC.

## Installation
```bash
cd server
protoc --go_out=. --go-grpc_out=. ../pb_schemas/currency.proto
```

## Run
```bash
cd server
go run server.go

# second terminal
cd client
node client.js
```