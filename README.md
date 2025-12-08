# Internal Transfers System (Go + Postgres)

## Overview
A simple internal transfers HTTP service that supports:
- Create account: `POST /accounts`
- Query account: `GET /accounts/{id}`
- Submit transfer: `POST /transactions`


## Requirements
- Docker & docker-compose OR
- Go 1.23+ and Postgres
- .env file
  - DB_USER={username}
  - DB_PASSWORD={password}
  - DB_NAME={db_name}
  - DB_PORT={db_port}

## Running with Docker Compose
```bash
docker-compose up --build
```

## Example curl commands

### Create accounts:

``` bash
curl -X POST http://localhost:8080/api/v1/accounts \
  -H 'Content-Type: application/json' \
  -d '{"account_id":123,"initial_balance":"100.23344"}'
```

``` bash
curl -X POST http://localhost:8080/api/v1/accounts \
  -H 'Content-Type: application/json' \
  -d '{"account_id":456,"initial_balance":"0.1"}'
```


### Get account:

``` bash
curl http://localhost:8080/api/v1/accounts/123
```

### Create transaction:

``` bash
curl -X POST http://localhost:8080/api/v1/transactions \
  -H 'Content-Type: application/json' \
  -d '{"source_account_id":123,"destination_account_id":456,"amount":"50.12345"}'
```
