run:
	go run cmd/main/*.go

migrate:
	go run cmd/main/*.go migrate

run-db:
	cd build/docker; docker-compose up -d postgres
