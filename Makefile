DB_NAME=chatroom
DB_USER=myuser
DB_PASS=mypassword
DB_HOST=localhost
DB_PORT=5432
DATABASE_DSN ?= postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable
MIGRATIONS_DIR = ./db/migrations

# So psql doesn't prompt for password
export PGPASSWORD=$(DB_PASS)

db-create:
	psql -U $(DB_USER) -h $(DB_HOST) -p $(DB_PORT) -c "CREATE DATABASE $(DB_NAME);"

db-drop:
	psql -U $(DB_USER) -h $(DB_HOST) -p $(DB_PORT) -c "DROP DATABASE IF EXISTS $(DB_NAME);"

db-reset: db-drop db-create

db-migrate:
	goose -dir ${MIGRATIONS_DIR} postgres "${DATABASE_DSN}" up

db-migrate-down:
	goose -dir ${MIGRATIONS_DIR} postgres "${DATABASE_DSN}" down

db-create-migration:
ifndef name
	$(error You must provide a migration name: make migration name=create_table)
endif
	goose -dir $(MIGRATIONS_DIR) create $(name) sql
