ProjectName := github.com/SomethingBot/discordingestor
GitCommit   := $(shell git rev-parse HEAD)
GitTag      := $(shell git tag --points-at HEAD)

build:
	go build -ldflags "-X main.GitCommit=$(GitCommit) -X main.GitTag=$(GitTag) -X main.Mode=Dev"

test:
	go vet -race ./...
	go test -short ./...

spawn-redis:
	podman run -p 6379:6379 --rm --name alias-redis-dev docker.io/library/redis

test-integration:
	go test

lint:
	go fmt ./...

run: build
	DISCORD_APIKEYFILE=apikeyfile REDIS_ENDPOINTS=127.0.0.1:6379 ./discordingestor
