#!/bin/bash

printf "\033c"
printf "\033[1;33mCompiling ...\033[0m\n"

BIN_PATH="$(git rev-parse --show-toplevel)/prototype/.out/server.out"

function start_server() {
    $BIN_PATH &
}

function stop_server() {
    PID=$(ps -ef | grep -v grep | grep $BIN_PATH | awk '{print $2}')
    kill $PID >>/dev/null
}

if go build -o $BIN_PATH server/main.go; then

    printf "\033c"
    printf "\033[1;32mCompiled successfully!\n\033[1;33mServer running 😋 ...\033[0m\n"

    # stop_server >>/dev/null
    stop_server
    start_server
    
else
    printf "\033[1;31mCompilation failed!\n😟\033[0m\n"
fi
