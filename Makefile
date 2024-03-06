install:
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
templ: 
	@templ generate -watch -proxy=http://localhost:1323
tailwind:
	@tailwindcss -i ./assets/css/styles.css -o ./public/tailwind.css --watch
migrate-create:
	echo "Creating migration wywy_db"
	migrate create -ext sql -dir db/migrations -seq wymy_db
migrate-up:
	echo "Migrating-up database"
	migrate -source file://db/migrations -database 'postgres://wywy:123456@localhost:4444/wywy_db?sslmode=disable' -verbose up
migrate-down:
	echo "Migrating-down database"
	migrate -source file://db/migrations -database 'postgres://wywy:123456@localhost:4444/wywy_db?sslmode=disable' -verbose down
