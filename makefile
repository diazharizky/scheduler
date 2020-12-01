run:
	go run cmd/main/*.go

migrate-up:
	go run cmd/main/*.go migrate-up

migrate-down:
	go run cmd/main/*.go migrate-down

run-db:
	cd build/docker; docker-compose up -d postgres
