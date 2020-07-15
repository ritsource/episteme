#!/bin/bash

BIN_PATH="$(git rev-parse --show-toplevel)/prototype/.out/server.out"

function stop_server() {
    PID=$(ps -ef | grep -v grep | grep $BIN_PATH | awk '{print $2}')
    if kill -0 $PID; then
        kill $PID >>/dev/null
    fi
    # kill $PID >>/dev/null
}

# source - https://stackoverflow.com/questions/360201/how-do-i-kill-background-processes-jobs-when-my-shell-script-exits
trap "stop_server" EXIT

# trap "stop_server; printf \"\033[1;33mServer stopped 😃!\033[0m\n\"; exit" 2
# trap "stop_server; echo \"Hello world\"" EXIT
# trap "stop_server" SIGINT SIGTERM EXIT

# variable for keeping track of build state
# there's no good in trigger a compile while
# there's alread a compilation going on
export BUILD_IN_PROCESS=0

# while compiling `run-server.sh` ignore all
# the file change events. this variable tells
# the script if there has been a latest change
# is the source code, so it can recompile once
# the current compilation is done
export THERE_EXISTS_A_NEWER_CHANGE=0

# source - https://stackoverflow.com/a/38229197/9406420
# ls **/*.go | entr bash scripts/run-server.sh
find . -name "*.go" | entr bash scripts/run-server.sh
