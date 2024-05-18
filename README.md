# wywy
Go + HTMX
## How to run

    air -c .air.dev.toml



## How to migration

    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

    make migrate-up
