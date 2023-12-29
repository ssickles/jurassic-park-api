# Jurassic Park API
An API to control the cages and dinosaurs in Jurassic Park.

## Getting Started
To run this application locally, first clone this repository and make sure you have [Docker](https://www.docker.com/) installed.

Then, run the following command in the root directory of the project:
```shell
docker-compose up --build
```

This will stand up an instance of Postgres as well as this API in docker containers.
The API is pre-configured with a connection to the Postgres db via the `POSTGRES_URL` environment variable in the `docker-compose.yml` file.

This is all that is required to run this application. The API will be hosted at http://localhost:8888.

## Migrations
Database migration management has been added using the [goose](https://github.com/pressly/goose) library.
All migrations are located in the /migrations directory are configured to run when the docker container starts up.

The Dockerfile has been configured with an entrypoint script to run migrations before starting the API.
This way, all migrations must be run before a new version of the API is run.

[See the `/docker-entrypoint.sh` file for the goose command]

### Manual Migrations
To run migrations manually, you must:
1. [Install goose](https://pressly.github.io/goose/installation/) on your local machine
2. Make sure the Postgres db is running locally
```shell
docker-compose up db
```

Then, run your goose command to apply, remove, or check the status of migrations.
Below is an example:
```shell
export POSTGRES_URL="postgres://jurassic-park:p@$$w0rd@db:5432/jurassic-park?sslmode=disable"
goose -dir=/app/migrations/ postgres "${POSTGRES_URL}" up
```

## GIN Web Framework
This API is built using the [Gin Web Framework](https://gin-gonic.com/).
The Gin framework provides a performant and flexible web framework to reduce boilerplate code and help provide a consistent way of developing our APIs.

## Models and Data Layer
All data models are located in the /models directory.
These models have decorators for serialization by the data layer to persist data to the database.
The models also have decorators to serialize the data for API responses.

## API Documentation
All routes for this api are defined in the /api directory and configured in setup_router.go.
