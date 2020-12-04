.PHONY: build clean clean-packr

run:
	source .env; go run cmd/scheduler/*.go

migrate-up:
	go run cmd/scheduler/*.go migrate-up

migrate-down:
	go run cmd/scheduler/*.go migrate-down

run-db:
	cd build/docker; docker-compose up -d postgres pgadmin

generate:
	go get -v ./... && go generate -v ./...

build:
	for dir in `find ./cmd -name main.go -type f`; do \
		go build -v -o "bin/$$(basename $$(dirname $$dir))" "$$(dirname $$dir)"; \
	done

clean:
	rm -rf bin

clean-packr:
	cd cmd/scheduler &&\
	go run github.com/gobuffalo/packr/v2/packr2 clean