FROM golang:alpine

RUN apk update && apk add --no-cache git

ENV PORT 8080

WORKDIR /go/src/github.com/ritsource/episteme

COPY . .

RUN cd ./prototype && go get ./...

RUN git rev-parse --show-toplevel

RUN go run ./prototype/tools/data_encoder/yml/main.go

RUN mkdir -p /app/.out

RUN cd prototype/server && go build -o /app/.out/server-prod

CMD [ "/app/.out/server-prod" ]
