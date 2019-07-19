# Migrations

Run
```
go get github.com/codegangsta/gin
gin
go to http://localhost:3000 
```

Install migrate tool
```
go get -tags 'mysql' -u github.com/golang-migrate/migrate/cmd/migrate
```

New migrations
```
migrate create -ext mysql -dir migrations migration_name
```

Migrate up
```
migrate -path ./migrations -database="mysql://root:pass@tcp(127.0.0.1:3306)/dbname" up
migrate -path ./migrations -database="mysql://root:pass@tcp(127.0.0.1:3306)/dbname" down
```