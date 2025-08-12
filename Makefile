DB_NAME=chatroom
DB_USER=myuser
DB_PASS=mypassword
DB_HOST=localhost
DB_PORT=5432

# So psql doesn't prompt for password
export PGPASSWORD=$(DB_PASS)

db-create:
	psql -U $(DB_USER) -h $(DB_HOST) -p $(DB_PORT) -c "CREATE DATABASE $(DB_NAME);"

db-drop:
	psql -U $(DB_USER) -h $(DB_HOST) -p $(DB_PORT) -c "DROP DATABASE IF EXISTS $(DB_NAME);"

db-reset: db-drop db-create

