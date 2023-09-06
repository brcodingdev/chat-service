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

## Setup and run
 - Update .env file
## Run

```bash
# run the service
$ make run
```

### Test

```bash
# run tests
$ make test
```


## TODO
 - More Unit Tests
 - Run inside docker container
 - Improve the code with the best practices in Golang