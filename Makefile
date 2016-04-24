.PHONY: server, deps, install_server

GO_APP_PATH=${GOPATH}/src/github.com/lean-poker

init: mkdir_app link deps install_server

link:
	rm -f ${GO_APP_PATH}/poker-player-go && ln -fs ${PWD} ${GO_APP_PATH}/poker-player-go

mkdir_app:
	mkdir -p ${GO_APP_PATH}

install_server:
	go install -a player-service.go

server: deps
	go run player-service.go
deps:
	go get github.com/lean-poker/poker-player-go/player && \
	go get github.com/lean-poker/poker-player-go/leanpoker
