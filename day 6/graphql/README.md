# graphql

![golangci-lint status](https://github.com/103cuong/graphql/workflows/golangci-lint/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/103cuong/graphql)](https://goreportcard.com/report/github.com/103cuong/graphql)

Building 🍣 a GraphQL Server with Go.

## commands

```shell script
go get github.com/99designs/gqlgen@v0.11.3
go run github.com/99designs/gqlgen init
go run github.com/99designs/gqlgen generate
docker run --name mysql -e MYSQL_ROOT_PASSWORD=test -e MYSQL_DATABASE=test -d mysql:latest
```

## installation

```shell script
go mod download
```

## usage

```shell script
go run main.go
```

*go to http://localhost:8080/*

## documents

- [gqlgen](https://gqlgen.com/)
- [graphql-go Tutorial](https://www.howtographql.com/graphql-go/0-introduction/)

## license

MIT © [Cuong Tran](https://github.com/103cuong)
