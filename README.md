# CRUD Golang Project Documentation

This documentation provides an overview of the CRUD Golang project. The project is designed to perform CRUD (Create, Read, Update, Delete) operations on a user entity using a MySQL database.

## Project Structure

The project consists of the following files and directories:

- `main.go`: The main entry point of the application, which sets up the HTTP server and routes.
- `go.mod`: The Go module file that manages project dependencies.
- `database/`: A package containing functions for establishing a database connection.
- `services/userService.go`: A package containing user-related functions for handling CRUD operations.

## Dependencies

The project relies on the following third-party packages, which are managed by Go modules:

- `github.com/go-sql-driver/mysql`: A MySQL driver for Go.
- `github.com/gorilla/mux`: A powerful HTTP router and URL matcher for building web applications.

## Installation

To run the project locally, follow these steps:

1. Make sure you have Go installed on your machine.
2. Clone the project repository.
3. Navigate to the project directory in your terminal.
4. Run the following command to install the project dependencies:

```shell
go mod download
```

5. Once the dependencies are downloaded, run the following command to start the application:

```shell
go run main.go
```

6. The server will start running on `http://localhost:5000`. You can now access the different routes to perform CRUD operations on users.

## API Endpoints

The project exposes the following API endpoints:

- `POST /users`: Creates a new user. Requires a JSON payload with the user's name and email.
- `GET /users`: Retrieves all users from the database.
- `GET /users/{id}`: Retrieves a user by ID from the database.
- `PUT /users/{id}`: Updates a user's data in the database. Requires a JSON payload with the updated user information.
- `DELETE /users/{id}`: Deletes a user from the database.

## Response Format

The API endpoints generally return JSON responses with appropriate HTTP status codes. Successful requests return the requested data, while errors are accompanied by error messages.

## Error Handling

The project implements a generic error handling mechanism using the `handleGenericError` function. It captures and handles common errors encountered during database operations or request processing.

## Conclusion

The CRUD Golang project provides a simple and efficient way to perform CRUD operations on user data stored in a MySQL database. By following the installation instructions and using the provided API endpoints, you can easily create, read, update, and delete user records.

For more details on each function and implementation, refer to the source code in the respective files.

If you have any further questions or need assistance, please feel free to ask.

**Note:** Replace `localhost:5000` with the appropriate host and port if you deploy the project to a different environment.