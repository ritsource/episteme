#!/bin/bash
set -e

# get sha hash for unieque container-id
GIT_SHA=$(git rev-parse HEAD)

# google cloud project id (as command line arg)
PROJECT_ID=$1

if [[ -z "$PROJECT_ID" ]]; then
  # exit if PROJECT_ID not provided
  echo "PROJECT_ID not provided"
  exit 1
else
  :
fi

# testing go server
# go test ./...

# building docker container
docker build -t ritsource/episteme .

# pushing to google container registary
docker tag ritsource/episteme gcr.io/$PROJECT_ID/ritsource-episteme-prototype-server:latest
docker tag ritsource/episteme gcr.io/$PROJECT_ID/ritsource-episteme-prototype-server:$GIT_SHA

# pushing to google container registary
docker push gcr.io/$PROJECT_ID/ritsource-episteme-prototype-server:latest
docker push gcr.io/$PROJECT_ID/ritsource-episteme-prototype-server:$GIT_SHA
