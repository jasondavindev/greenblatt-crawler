.PHONY: build
build:
	go build -o build/greenblatt_crawler

dkbuild:
	docker run --rm -v ${PWD}:/app -w /app golang bash -c "go build -o build/greenblatt_crawler"
