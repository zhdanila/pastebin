# Pastebin Application README

## Dependencies

To run the Pastebin application, ensure you have the following dependencies installed:

- [github.com/google/logger](https://github.com/google/logger) - Logging library
- [github.com/jmoiron/sqlx](https://github.com/jmoiron/sqlx) - Library for PostgreSQL
- [github.com/lib/pq](https://github.com/lib/pq) - PostgreSQL driver
- [github.com/spf13/viper](https://github.com/spf13/viper) - Key-value management tool
- [github.com/joho/godotenv](https://github.com/joho/godotenv) - Tool for storing essential data
- [github.com/go-redis/redis/v8](https://github.com/go-redis/redis) - Redis library
- [github.com/aws/aws-sdk-go](https://github.com/aws/aws-sdk-go) - AWS SDK for Go

## Configuration

1. Replace placeholders in `config.yml` and `.env` with your Amazon S3 credentials.
2. Initialize PostgreSQL using Docker:
   ```
   sudo docker run --name=pastebin -e POSTGRES_PASSWORD=qwerty -p 5436:5432 -d postgres
   ```
3. Run Redis using Docker:
   ```
   sudo docker run -d --name redis_pastebin -p 6380:6379 redis
   ```

## Running Migrations

Apply migrations to PostgreSQL:
```
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
```

## Endpoints

- **POST /pastebin/**: Send your text to store.
  ```
  {
    "text": "Your text",
    "password": "Optional",
    "pasteTTL": "Period for storing the text"
  }
  ```

- **GET /pastebin/{id}**: Retrieve text by ID.
