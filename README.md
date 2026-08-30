## BankApi (TakeHome)

This repo provides basic functionality for tracking **expenses** written in Go 1.27

Stack used: Chi, sqlx, swaggo, PostgreSQL

## Q: Why to pick this kind of stack ?
A: It's equivalent to Rust (Axum + sqlx) stack, providing an access to raw SQL Queries. 

## Quick Start
First of all - having Go 1.27 is required on your machine.

Create **.env** file in root directory and add following code:

```shell
PRODUCTION=N
DATABASE_URL=postgresql://username:password@host:port/BankDb?sslmode=disable
DATABASE_TEST_URL=postgresql://username:password@host:port/BankTestDb?sslmode=disable
```

For manual usage:
```shell
git clone https://github.com/peidxxk/Bank.git
go install
swag init -g main.go
go run .
```

For testing purposes:
```shell
go test ./test -v
```
