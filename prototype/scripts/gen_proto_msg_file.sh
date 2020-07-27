#!/bin/bash

BASEDIR=$(dirname "$0")
# echo "$BASEDIR"

DEFAULT_INPUT_PATH="$BASEDIR/../.data/posts.yml"
DEFAULT_OUTPUT_PATH="$BASEDIR/../.data/dst/prod"

# go run tools/data_encoder/yml/main.go .data/posts.yml .data/dst/prod
go run "$BASEDIR/../tools/data_encoder/yml/main.go" $DEFAULT_INPUT_PATH $DEFAULT_OUTPUT_PATH
