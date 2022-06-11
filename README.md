## Tsarka Task

## Lauch

```
$ docker-compose up
$ migrate -path ./schema -database 'postgres://postgres:123456Aa@localhost:5432/postgres?sslmode=disable' up
$ go run cmd/main.go 
```

file '.env' must be deleted...