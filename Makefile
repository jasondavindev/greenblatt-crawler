.PHONY: build
build:
	go build -o build/statusinvest_crawler

dkbuild:
	docker run --rm -v ${PWD}:/app -w /app golang bash -c "go build -o build/statusinvest_crawler"
