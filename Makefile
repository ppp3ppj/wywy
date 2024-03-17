# Define variables
CSS_DIR := assets/css
PUBLIC_DIR := public
TAILWIND_CSS := $(CSS_DIR)/styles.css
OUTPUT_CSS := $(PUBLIC_DIR)/tailwind.css
PORT := 1323

install:
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

templ: 
	@templ generate -watch -proxy=http://localhost:1323

migrate-create:
	echo "Creating migration wywy_db"
	migrate create -ext sql -dir db/migrations -seq wymy_db

migrate-up:
	echo "Migrating-up database"
	migrate -source file://db/migrations -database 'postgres://wywy:123456@localhost:4444/wywy_db?sslmode=disable' -verbose up

migrate-down:
	echo "Migrating-down database"
	migrate -source file://db/migrations -database 'postgres://wywy:123456@localhost:4444/wywy_db?sslmode=disable' -verbose down

# tailwind build
# Default target
.PHONY: all
all: build

# Build CSS
.PHONY: build
build: $(OUTPUT_CSS)

$(OUTPUT_CSS): $(TAILWIND_CSS)
	@mkdir -p $(PUBLIC_DIR)
	npx tailwindcss -i $(TAILWIND_CSS) -o $(OUTPUT_CSS)

# Watch Tailwind CSS
.PHONY: watch
watch:
	npx tailwindcss -i $(TAILWIND_CSS) -o $(OUTPUT_CSS) --watch

# Resolve port and kill process if exists
.PHONY: resolve-port
resolve-port:
	@echo "Checking if port $(PORT) is in use..."
	@if lsof -i :$(PORT); then \
		echo "Port $(PORT) is in use, killing the process..."; \
		kill -9 `lsof -t -i :$(PORT)`; \
	else \
		echo "Port $(PORT) is not in use."; \
	fi

# Clean CSS
.PHONY: clean
clean:
	@rm -f $(OUTPUT_CSS)

