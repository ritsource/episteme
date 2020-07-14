#!/bin/bash

BIN_PATH="$(git rev-parse --show-toplevel)/prototype/.out/server.out"

function start_server() {
    $BIN_PATH &
}

function stop_server() {
    PID=$(ps -ef | grep -v grep | grep $BIN_PATH | awk '{print $2}')
    kill $PID >>/dev/null
    # kill $PID
    if [ $? -eq 0 ]; then
        printf "Server stopped ...."
    else
        printf "\033[1;31mUnable to kill process! PID=$PID\n😟\033[0m\n"
    fi
}

if [ $BUILD_IN_PROCESS -eq 0 ]; then
    # printf "\033c"
    printf "\033[1;33mCompiling ...\033[0m\n"

    go build -o $BIN_PATH server/main.go

    if [ $THERE_EXISTS_A_NEWER_CHANGE -eq 1 ]; then
        # printf "\033c"
        printf "\033[1;33mCompiling ...\033[0m\n"

        go build -o $BIN_PATH server/main.go
    fi

    BUILD_IN_PROCESS=0
    THERE_EXISTS_A_NEWER_CHANGE=0
    
    if [ $? -eq 0 ]; then
        # printf "\033c"
        printf "\033[1;32mCompiled successfully!\n\033[1;33mServer running 😋 ...\033[0m\n"

        if kill -0 $(ps -ef | grep -v grep | grep $BIN_PATH | awk '{print $2}'); then
            stop_server
        fi
        # stop_server >>/dev/null
        start_server

    else
        printf "\033[1;31mCompilation failed!\n😟\033[0m\n"
    fi
else
    THERE_EXISTS_A_NEWER_CHANGE=1
fi   
