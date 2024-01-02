# Jurassic Park API
An API to control the cages and dinosaurs in Jurassic Park.

## Running Locally via Docker
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

## GIN Web Framework
This API is built using the [Gin Web Framework](https://gin-gonic.com/).
The Gin framework provides a performant and flexible web framework to reduce boilerplate code and help provide a consistent way of developing our APIs.

## App Structure
The application is structured into the following packages:
1. `api` - Contains all the routes and controllers for the API.
2. `park` - Contains all the business logic for the application.
3. `models` - Contains all the data models for the application.
4. `data` - Provides an abstraction layer for the database.

## Api Package
The `setup_router.go` file in the /api directory is where all the routes are defined.
This file refers to the controllers in the /api directory to handle the requests.
Any logic concerning validation of inputs and providing responses is handled in the controllers.
Any business logic is handled in the park package.

## Park Package
The `park` package contains all the business logic for the application.
It should not contain any logic concerning the API or the database.

## Models Package
All data models are located in the /models directory.
These models have decorators for serialization by the data layer to persist data to the database.
The models also have decorators to serialize the data for API responses.
The models provide the data structures for the application.

## Data Package
The `data` package provides an abstraction layer for the database.
This package contains all the database queries and logic to persist data to the database.
This layer could be swapped out for a different database implementation without affecting the rest of the application.

## TODO
The following items are some things that I would have liked to do given more time:
1. Continue to refine the `park` package as more functionality is added.
This package has started as collection of functions to contain the business logic.
As the application grows, this package will likely evolve into a struct or structs to contain the business logic.
This refinement is best done as the patterns reveal themselves.
2. Add more validation errors, provide better error messages, and add error codes.
I'm also sure there are some edge cases that I haven't covered.
3. Wrap the database calls in a transaction to ensure that all database calls are atomic.
4. Continue to add tests for all the API endpoints.
5. Add authorization so that only builders can create new cages and scientists can create, cage, and view the status of dinosaurs.

Bonus Points I did not get to:
1. Cages know how many dinosaurs are contained.
- The query exists to get the number of dinosaurs in a cage, but it hasn't been exposed in the API.
2. Cages cannot be powered off if they contain dinosaurs.
- I haven't built an endpoint to power off a cage, but the query exists to check if a cage contains dinosaurs.
3. When querying dinosaurs or cages they should be filterable on their attributes (Cages on their power status and dinosaurs on species).
