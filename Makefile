ProjectName := github.com/SomethingBot/discordingestor
GitCommit   := $(shell git rev-parse HEAD)
GitTag      := $(shell git tag --points-at HEAD)

build:
        go build -ldflags "-X main.GitCommit=$(GitCommit) -X main.GitTag=$(GitTag) -X main.Mode=Dev"

test:
        go vet ./...
        go test ./...

spawn-redis:
        podman run -p 6379:6379 --rm --name alias-redis-dev docker.io/library/redis

test-integration:
        go test -tags=integration

lint:
        go fmt ./...

run: build
        DISCORDINGESTOR_REDIS_ENDPOINTS=localhost:2379 GATEWAY_ADDRESS=0.0.0.0:8987 ./discordingestor