.PHONY: wiregen up

wiregen:
	third_party/bin/wire gen ./...

up:
	go run -race ./cmd/server
