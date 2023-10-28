include .env

create_container:
	docker run --name ${DB_CONTAINER} -p ${DB_PORT}:5432 -e POSTGRES_USER=${USER} -e POSTGRES_PASSWORD=${PASSWORD} -d postgres:15-alpine

postgres:
	docker exec -it ${DB_CONTAINER} psql

create_db:
	docker exec -it ${DB_CONTAINER} createdb --username=${USER} --owner=${USER} ${DB_NAME}
	
dropdb:
	docker exec -it ${DB_CONTAINER} dropdb projects-db

start_containers:
	docker start ${DB_CONTAINER}

create_migrations:
	sqlx migrate add -r init

migrate_up:
	sqlx migrate run --database-url "postgres://${USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

migrate_down:
	sqlx migrate revert --database-url "postgres://${USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

run: create_container create_db migrate_up
	@echo "starting api"