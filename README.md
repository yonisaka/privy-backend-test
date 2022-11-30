# privy-backend-test 🔥
Pretest PrivyID Backend Engineer
Menggunakan Golang, Mysql sebagai database.

## Framework dan Scaffolding

### [GIN](https://github.com/gin-gonic/gin)
Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity.

### [Testify](https://github.com/stretchr/testify)
Go code (golang) set of packages that provide many tools for testifying that your code will behave as you intend.

### [Go Validator](https://github.com/go-playground/validator)
Package validator implements value validations for structs and individual fields based on tags.

### [GoDotEnv](https://github.com/joho/godotenv)
A Go (golang) port of the Ruby dotenv project (which loads env vars from a .env file).

### [Goose](https://github.com/pressly/goose)
Goose is a database migration tool. Manage your database schema by creating incremental SQL changes or Go functions.

### [Uuid](https://github.com/gofrs/uuid)
Package uuid provides a pure Go implementation of Universally Unique Identifiers (UUID) variant as defined in RFC-4122. This package supports both the creation and parsing of UUIDs in different formats.

### [Jwt-go](https://github.com/dgrijalva/jwt-go)
In short, it's a signed JSON object that does something useful (for example, authentication). It's commonly used for Bearer tokens in Oauth 2. A token is made of three parts, separated by .'s. The first two parts are JSON objects, that have been base64url encoded. The last part is the signature, encoded the same way.

### [Sqlx](https://github.com/jmoiron/sqlx)
Sqlx is a library which provides a set of extensions on go's standard database/sql library.

### [Go-SqlMock](https://pkg.go.dev/gopkg.in/DATA-DOG/go-sqlmock.v1#section-readme)
Sqlmock is a mock library implementing sql/driver. Which has one and only purpose - to simulate any sql driver behavior in tests, without needing a real database connection. It helps to maintain correct TDD workflow.

### [Logrus](https://github.com/sirupsen/logrus)
Logrus is a structured logger for Go (golang), completely API compatible with the standard library logger.

## Requirements
Membutuhkan requirements sebagai berikut.

| Requirement | Version |
| ----------- | ----------- |
| Go | >= 1.17.2 |
| Mysql | = 5.7.33 |

## Installation
Pastikan semua requirements telah di terinstall di sistem.
Clone projek ini dan install dependency yang dibutuhkan

```
$ git clone https://github.com/yonisaka/privy-backend-test
$ cd privy-backend-test
$ go mod tidy
```

Selanjutnya membuat database **privy_backend_test** di lokal mysql

## Configuration
Copy **.env.default** menjadi **.env** 

## Migration & Seeder
Migrate dan Seed menjadi satu file eksekusi:

Command Up:
```
$ go run database/migrate.go up
```

Command Down:
```
$ go run database/migrate.go down
```
## Run Application
Pastikan Mysql dan Redis sudah berjalan.

Run Application :
```
$ go run cmd/main.go
```

Atau dapat menggunakan air :
```
$ air
```

## Test
Run Unit Test Usecase: 
```
$ go test internal/repositories/cake_repository_test.go -v
$ go test internal/repositories/user_repository_test.go -v
```

Run Unit Test Repository: 
```
$ go test internal/usecases/cake_usecase_test.go -v
$ go test internal/usecases/user_usecase_test.go -v
```

## Docker
Copy **.env.docker.default** menjadi **.env** 
Untuk menjalankan menggunakan docker, bisa menggunakan docker compose
```
$ docker-compose up -d
```

## Postman Collection
Untuk postman collection dapat di import dari **Privy Test.postman_collection**

### Postman Env
| Key | Value |
| ----------- | ----------- |
| url | localhost:8080 |
| bearer | **generated token** |