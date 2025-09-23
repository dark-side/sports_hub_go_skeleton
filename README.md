# Sports-Hub Application Go Back-End

## Project Description

This is a draft pet project for testing Generative AI on different software engineering tasks. It is planned to evolve and grow over time. Specifically, this repo will be a Go (Golang) playground with Gin framework, GORM ORM, and Docker Compose. The application's legend is based on the sports-hub application description from the following repo: [Sports-Hub](https://github.com/dark-side/sports-hub).

## Available Front-End applications
- [React.js](https://github.com/dark-side/sports_hub_react_skeleton)
- [Angular](https://github.com/dark-side/sports_hub_angular_skeleton)

## Dependencies

- Docker
- Docker Compose

The mentioned dependencies can be installed using the official documentation [here](https://docs.docker.com/compose/install/).
[Podman](https://podman-desktop.io/docs/compose) can be used as an alternative to Docker.

## Setup and Running the Application

### Clone the Repositories

To run the web application with the React front-end, clone the following repositories within the same folder:

```sh
git clone git@github.com:dark-side/sports_hub_go_skeleton.git
git clone git@github.com:dark-side/sports_hub_angular_skeleton.git
```

### Navigate to the back-end application directory

All commands should be run from the `sports_hub_go_skeleton` directory.

### Copy environment variables

Before running the project, copy the example environment file:

```bash
cp .example.env .env
```

### Build and run the project
```bash
make run
```
This command will build the Docker images and start the containers defined in the `docker-compose.yml` file.

For Windows you might need to install [Make](https://www.gnu.org/software/make/#download) first or run the command (`-d` for detached mode to run in the background):
```bash
docker compose up --build -d
```

### Access the application
Once the containers are up and running, you can access the application at:
- **Frontend**: [http://localhost:3000](http://localhost:3000)
- **Backend**: [http://localhost:3002](http://localhost:3002)

### API Endpoints

#### Register
•	Endpoint: POST http://localhost:3002/api/auth/register
•	Description: Register a new user.
•	Request Body Example:
```json
{
  "name": "test",
  "email": "test@gmail.com",
  "password": "test"
}
```

#### Login
* Endpoint: POST http://localhost:3002/api/auth/login
* Description: Authenticate user and get JWT token.
* Request Body Example:
```json
{
  "email": "test@gmail.com",
  "password": "test"
}
```

#### User

Get Profile
* Endpoint: GET http://localhost:3002/api/user/profile
* Description: Retrieve authenticated user’s profile.
* Authorization: Required — Add Header - Authorization:  <your_token>.

#### Articles

Get All Articles
* Endpoint: GET http://localhost:3002/api/articles
* Description: Retrieve all articles.
* Authorization: Required — Add Header - Authorization:  <your_token>.

#### Create Article
* Endpoint: POST http://localhost:3002/api/articles
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

#### Update Article
* Endpoint: PUT http://localhost:3002/api/articles/{id}
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

#### Delete Article
*   Endpoint: DELETE http://localhost:3002/api/articles/{id}
*   Description: Delete an article by ID.
*   Authorization: Required — Add Header - Authorization:  <your_token>.

#### Get Article by ID
*   Endpoint: GET http://localhost:3002/api/articles/{id}
*   Description: Retrieve a single article by ID.
*   Authorization: Required — Add Header - Authorization:  <your_token>.

## License

Licensed under either of

- [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)
- [MIT license](http://opensource.org/licenses/MIT)

Just to let you know, at your option.

## Contribution
Unless you explicitly state otherwise, any contribution intentionally submitted for inclusion in your work, as defined in the Apache-2.0 license, shall be dual licensed as above, without any additional terms or conditions.

**Should you have any suggestions, please create an Issue for this repository**
