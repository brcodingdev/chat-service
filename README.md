# Chat Service

Chat Service is a messaging application that enables users to send messages in chat rooms. It leverages technologies like Golang, PostgreSQL, JWT authentication, and ORM with gorm.

## Technologies
- WebSockets
- PostgreSQL
- golang-jwt
- gorm
- testify


## Requirement
- Docker
- docker-compose
- Golang `>=` 1.20
- golangci-lint


## Related services
- stock-service : Service that calls an API using a "stock_code".  <br/>
  <b>Repo : https://github.com/brcodingdev/stock-service.git </b>


- chat-frontend: The frontend application that allows run commands. <br/>
  <b>Repo: https://github.com/brcodingdev/chat-frontend.git </b>

## Setup
Update .env file, env vars:

```
#PostgreSQL

DB_HOST=localhost
DB_PORT=5432
DB_DATABASE=chat
DB_USERNAME=postgres
DB_PASSWORD=postgres
 
#RabbitMQ  

RABBIT_USERNAME=guest
RABBIT_PASSWORD=guest
RABBIT_HOST=localhost
```

## Run

### build docker image (optional)

```bash
# builds an image
$ make build-docker
```

### run outside docker (optional)

```bash
# run the service
$ make run
```

### run with docker

```bash
# run inside container
$ make run-docker
```

### Test

```bash
# run tests
$ make test
```

## TODO
 - More Unit Tests
 - Improve the code with the best practices in Golang