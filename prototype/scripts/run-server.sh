#!/bin/bash

printf "\033c"

printf "\033[1;33mCompiling ...\033[0m\n"

if go run server/main.go; then
    printf "\033c"

    printf "\033[1;32mCompiled successfully!\n\033[1;33mServer running 😋 ...\033[0m\n"
else
    printf "\033[1;31mCompilation failed!\n😟\033[0m\n"
fi
