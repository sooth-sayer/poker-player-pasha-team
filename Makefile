.PHONY: server, deps, install_server

install_server:
	go install -a player-service.go

server:
	go run player-service.go
deps:
	go get github.com/lean-poker/poker-player-go/player && \
	go get github.com/lean-poker/poker-player-go/leanpoker
