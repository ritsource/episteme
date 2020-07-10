#!/bin/bash

if go run main.go; then
    printf "\033[1;32mcompiled successfully
\033[1;33mserver running 😋\033[0m\n"
else
    printf "\033[1;31mcompilation failed 😟\033[0m\n"
fi
