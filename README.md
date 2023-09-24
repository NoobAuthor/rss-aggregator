# RSS Aggregator

This is a simple RSS Aggregator written in Go using go-chi and godotenv.

## Getting Started

To get a copy of the project up and running on your local machine, follow these steps:

1. Clone the repository: `git clone https://github.com/NoobAuthor/rss-aggregator.git`
2. Install the dependencies: `go mod download`
3. Create a `.env` file and set the required environment variables.
4. Run the application: `go run main.go`

### Prerequisites

- Go 1.17 or later
- A PostgreSQL database

### Environment Variables

The following environment variables are required to run the application:

- `DATABASE_URL`: the URL of the PostgreSQL database
- `PORT`: the port number to listen on

### Installing

To install the application, run `go install` in the root directory of the project.

## Running the tests

To run the automated tests for this project, run `go test ./...` in the root directory of the project.

## Deployment

To deploy the application to a live system, you can use a tool like Docker or Kubernetes.

## Built With

- [go-chi](https://github.com/go-chi/chi) - A lightweight, idiomatic and composable router for building Go HTTP services.
- [godotenv](https://github.com/joho/godotenv) - A Go port of Ruby's dotenv library (Loads environment variables from `.env`).

## What You Can Learn

By building this project, I learned:

- How to use go-chi to build a RESTful API in Go.
- How to use godotenv to load environment variables from a `.env` file.
- How to use sqlx to interact with a PostgreSQL database in Go.
- How to write automated tests for a Go application.
- How to deploy a Go application to a live system using Docker or Kubernetes.
