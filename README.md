# Go BE Playground

A backend playground project built with Go (Golang), Gin framework, GORM ORM, and Docker Compose.

## Getting Started

### 1. Copy environment variables

Before running the project, copy the example environment file:

```bash
cp .example.env .env
```

### 2. Build and run the project
```bash
make run
```
This command will build the Docker images and start the containers defined in the `docker-compose.yml` file.

### 3. Access the application
Once the containers are up and running, you can access the application at:
- **Backend**: [http://localhost:8080](http://localhost:8080)

### API Endpoints
Auth

Register
•	Endpoint: POST http://localhost:8080/api/auth/register
•	Description: Register a new user.
•	Request Body Example:
```json
{
  "name": "test",
  "email": "test@gmail.com",
  "password": "test"
}
```

## Login
* Endpoint: POST http://localhost:8080/api/auth/login
* Description: Authenticate user and get JWT token.
* Request Body Example:
```json
{
  "email": "test@gmail.com",
  "password": "test"
}
```

## User

Get Profile
* Endpoint: GET http://localhost:8080/api/user/profile
* Description: Retrieve authenticated user’s profile.
* Authorization: Required — Add Header - Authorization:  <your_token>.

## Articles

Get All Articles
* Endpoint: GET http://localhost:8080/api/articles
* Description: Retrieve all articles.
* Authorization: Required — Add Header - Authorization:  <your_token>.

## Create Article
* Endpoint: POST http://localhost:8080/api/articles
* Description: Create a new article.
* Authorization: Required — Add Header - Authorization:  <your_token>.
* Request Body Example:
```json
{
  "title": "test",
  "short_description": "test",
  "description": "test"
}
```

## Update Article
* Endpoint: PUT http://localhost:8080/api/articles/{id}
* Description: Update an existing article by ID.
* Authorization: Required — Add Header - Authorization:  <your_token>.
* Request Body Example:
```json
{
  "title": "updated title",
  "short_description": "updated short description",
  "description": "updated description"
}
```

## Delete Article
*   Endpoint: DELETE http://localhost:8080/api/articles/{id}
*   Description: Delete an article by ID.
*   Authorization: Required — Add Header - Authorization:  <your_token>.

## Get Article by ID
*   Endpoint: GET http://localhost:8080/api/articles/{id}
*   Description: Retrieve a single article by ID.
*   Authorization: Required — Add Header - Authorization:  <your_token>.
