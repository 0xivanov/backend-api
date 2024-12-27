# Car Maintenance API

This is a simple Car Maintenance API built with Go, which allows users to manage cars, garages, and maintenance records. The API provides endpoints for CRUD operations and integrates Swagger for API documentation.

## Prerequisites

Before running the application, ensure you have the following installed:

- Go (version 1.16 or later)
- Go modules enabled
- A working Go environment

## Getting Started

Follow these steps to set up and run the Car Maintenance API:

### 1. Clone the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/0xivanov/backend-api.git
cd backend-api
```

### 2. Install Dependencies

Navigate to the project directory and install the required dependencies:

```bash
go mod tidy
```

### 3. Generate Swagger Documentation

To generate the Swagger documentation, you need to install the `swag` tool if you haven't already:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Then, run the following command to generate the documentation:

```bash
swag init
```

### 4. Run the Application

You can run the application using the following command:

```bash
go run main.go
```

The server will start on port `8088`. You should see a log message indicating that the server is running:

```
Server starting on port 8088...
```

### 5. Access the API

You can access the API endpoints using a tool like Postman or cURL. The base URL for the API is:

```
http://localhost:8088
```

### 6. Access Swagger UI

To view the API documentation, navigate to the following URL in your web browser:

```
http://localhost:8088/swagger/index.html
```

This will open the Swagger UI, where you can explore the available endpoints and test them directly.