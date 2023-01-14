# Go JWT Example

This repository is an example of using JSON Web Tokens (JWT) for authentication and authorization in a Go (golang) web application. The project includes basic implementation of JWT creation, signing and verification, and integration with a simple web server to demonstrate how JWT can be used to secure endpoints. The project includes an example of how to create a JWT and sign it with a secret key, as well as how to verify and decode the JWT on the server side to extract the claims and user information. Additionally, the project includes examples of how to use JWT middleware to secure endpoints and restrict access to only authenticated users.

This repository can serve as a starting point for developers looking to implement JWT in their Go web projects, and also helps to understand how JWT works and how it can be used to secure a web application.

### Getting Started
To get started with the project, simply clone the repository to your local machine and navigate to the project directory.
```sh
git clone https://github.com/sandeepkv93/go-jwt-example.git
cd go-jwt-example
```

### Running the Project
To run the project, simply execute the following command:

```sh
go run main.go
```

### Project Structure
The project is structured in the following way:

- main.go: The entry point for the application. It sets up the routes and handlers for the API.
- handlers/: This directory contains the handlers for the different routes in the application.
- utils/: This directory contains utility functions that are used throughout the application.

### JWT Implementation
The JWT implementation in this project is relatively simple. It uses the jwt-go library to handle the creation and validation of JWTs.

The project has two main endpoints: /login and /secure. The /login endpoint is used to log a user in and generate a JWT. The /secure endpoint is used to test the JWT authentication.

When a user makes a request to the /login endpoint, the application checks if the provided username and password match the hardcoded values in the code. If they match, a JWT is generated with the user's information and returned to the client.

When a user makes a request to the /secure endpoint, the application checks for the presence of a JWT in the request headers. If a JWT is present, the application attempts to validate the token. If the token is valid, the request is processed
