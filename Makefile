.PHONY: server, deps

server:
	go run player-service.go
deps:
	go get github.com/lean-poker/poker-player-go/player && \
	go get github.com/lean-poker/poker-player-go/leanpoker
