.PHONY: build
build:
	go build -o bin/greenblatt_crawler

dkbuild:
	docker run --rm -v ${PWD}:/app -w /app golang bash -c "go build -o bin/greenblatt_crawler"

run:
	go run main.go
