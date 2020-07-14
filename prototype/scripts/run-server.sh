#!/bin/bash

# printf "\033c"
printf "\033[1;33mCompiling ...\033[0m\n"

BIN_PATH="$(git rev-parse --show-toplevel)/prototype/.out/server.out"
# PID=""

# trap "trap - SIGTERM && print \"Hello world\"" INT SIGINT SIGTERM EXIT
# trap ctrl_c INT

# trap "echo \"Hello world\"; exit" INT SIGINT SIGTERM EXIT

# function run_server() {
#     PID=$(ps -ef | grep -v grep | grep $BIN_PATH | awk '{print $2}')
#     # printf "\$PID = $PID"
#     kill $PID
# 
#     # trap "kill $PID; printf \"PID=$PID\n\"" EXIT
#     $BIN_PATH
# }

# for i in `seq 1 5`; do
#     sleep 1
#     echo -n "."
# done

# trap "echo \"$PID=$(ps -ef | grep -v grep | grep $BIN_PATH | awk '{print $2}')\""

function start_server() {
    $BIN_PATH &
}

function stop_server() {
    PID=$(ps -ef | grep -v grep | grep $BIN_PATH | awk '{print $2}')
    # # printf "\$PID = $PID"
    kill $PID
}

if go build -o $BIN_PATH server/main.go; then
    # kill $PID >>/dev/null
    # rm BIN_PATH >>/dev/null

    # PID=$(ps -ef | grep -v grep | grep $BIN_PATH | awk '{print $2}')
    # # # printf "\$PID = $PID"
    # kill $PID

    # printf "\033c"
    printf "\033[1;32mCompiled successfully!\n\033[1;33mServer running 😋 ...\033[0m\n"

    # trap "printf \"PID=$PID\n\"" EXIT
    # $BIN_PATH &

    stop_server
    start_server
    
    # echo $! | read $PID
    # $BIN_PATH

    # trap "printf \"PID=$PID\n\"" EXIT
    # sleep 10
else
    printf "\033[1;31mCompilation failed!\n😟\033[0m\n"
fi

# source - https://stackoverflow.com/questions/360201/how-do-i-kill-background-processes-jobs-when-my-shell-script-exits
# trap "print \"Hello world\"" SIGINT SIGTERM EXIT

# trap "echo \"Hello world\"; exit" INT SIGINT SIGTERM EXIT

# trap "echo \"$PID=$(ps -ef | grep -v grep | grep $BIN_PATH | awk '{print $2}')\"; exit" EXIT
# sleep 1024 &    # background command
# BGPID=$!
# sleep 1024

# while /bin/true; do
#    sleep 1
#    echo -n "."
# done

# trap "stop_server; printf \"PID=$PID\n\"" EXIT
# for i in `seq 1 20`; do
#     sleep 1
#     echo -n "."
# done

