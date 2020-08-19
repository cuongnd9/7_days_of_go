# graphql_go

![golangci-lint status](https://github.com/103cuong/graphql_go/workflows/golangci-lint/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/103cuong/graphql_go)](https://goreportcard.com/report/github.com/103cuong/graphql_go)

implement ⚛️ graphql using golang (https://github.com/graph-gophers/graphql-go)

## installation

```shell script
go mod download
```

## usage

```shell script
go run main.go
```

## testing

```shell script
curl -XPOST -d '{"query": "{ hello }"}' localhost:8080/graphql
// {"data":{"hello":"Hello, @103cuong"}}
```

## license

MIT © [Cuong Tran](https://github.com/103cuong)