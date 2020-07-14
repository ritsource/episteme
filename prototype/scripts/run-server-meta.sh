#!/bin/bash

BIN_PATH="$(git rev-parse --show-toplevel)/prototype/.out/server.out"

function stop_server() {
    PID=$(ps -ef | grep -v grep | grep $BIN_PATH | awk '{print $2}')
    kill $PID >>/dev/null
    # printf "\033[1;33mServer stopped 😃!\033[0m\n"
}

# source - https://stackoverflow.com/questions/360201/how-do-i-kill-background-processes-jobs-when-my-shell-script-exits
trap "stop_server" EXIT
# trap "stop_server; printf \"\033[1;33mServer stopped 😃!\033[0m\n\"; exit" 2
# trap "stop_server; echo \"Hello world\"" EXIT
# trap "stop_server" SIGINT SIGTERM EXIT

# source - https://stackoverflow.com/a/38229197/9406420
ls **/*.go | entr bash scripts/run-server.sh
