# privy-rest-api 🔥
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
$ git clone https://gitlab.com/yonisaka/privy-backend-test
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
```

Run Unit Test Repository: 
```
$ go test internal/usecases/cake_usecase_test.go -v
```

## Docker
Untuk menjalankan menggunakan docker, bisa menggunakan docker compose
```
$ docker-compose up -d
```

## Postman Collection
Untuk postman collection dapat di import dari **Privy Test.postman_collection**