# aws-rest (Gin Gonic + AWS Go SDK)

## Description
This is a simple RESTful API built with Gin Gonic and AWS Go SDK.

## Features
- List all S3 Buckets
- Authentication middleware protected endpoints

## Prerequisites
- Go 1.16 or later
- AWS CLI credentials configured in your environment

## Guide
### 1. Clone the repository
```bash
git clone https://github.com/InspectorGadget/aws-rest.git
cd aws-rest
```

### 2. Setup the Go environment
```bash
go get
go mod tidy
```

### 4. Run the application
```bash
go run .
```

### 5. Test the API
You can use Postman or curl to test the API endpoints or use the built-in HTTP API Testkit in `./api-requests` folder.