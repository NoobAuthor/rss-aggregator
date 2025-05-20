# RSS Aggregator

A simple RSS Aggregator written in Go using [go-chi](https://github.com/go-chi/chi) and [godotenv](https://github.com/joho/godotenv).

## Table of Contents

- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [API Endpoints](#api-endpoints)
- [Running the Tests](#running-the-tests)
- [Deployment](#deployment)
- [Built With](#built-with)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

To get a copy of the project up and running on your local machine:

1. **Clone the repository:**
   ```sh
   git clone https://github.com/NoobAuthor/rss-aggregator.git
   cd rss-aggregator
   ```
2. **Install dependencies:**
   ```sh
   go mod download
   ```
3. **Create a `.env` file** in the project root and set the required environment variables (see [Configuration](#configuration)).
4. **Run the application:**
   ```sh
   go run ./cmd/main
   ```

## Configuration

The following environment variables are required:

- `DATABASE_URL`: URL of your PostgreSQL database (e.g., `postgres://user:pass@localhost:5432/dbname?sslmode=disable`)
- `PORT`: Port number for the HTTP server (e.g., `8080`)

Example `.env` file:

```
DATABASE_URL=postgres://user:pass@localhost:5432/rss?sslmode=disable
PORT=8080
```

## API Endpoints

| Method | Path         | Description           |
| ------ | ------------ | --------------------- |
| GET    | `/v1/health` | Health check endpoint |
| GET    | `/v1/error`  | Test error response   |

Example health check:

```sh
curl http://localhost:8080/v1/health
```

## Running the Tests

To run automated tests:

```sh
go test ./...
```

## Deployment

You can deploy the application using Docker or Kubernetes.

### Docker Example

1. Build the Docker image:
   ```sh
   docker build -t rss-aggregator .
   ```
2. Run the container:
   ```sh
   docker run --env-file .env -p 8080:8080 rss-aggregator
   ```

## Built With

- [go-chi](https://github.com/go-chi/chi) - HTTP router
- [godotenv](https://github.com/joho/godotenv) - Environment variable loader
- [sqlx](https://github.com/jmoiron/sqlx) - SQL database library (planned/used for DB access)

## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License.
