#!/bin/bash

cd /go/src/github.com/samra10/chat-api-go

go build

/usr/bin/wait-for-it.sh chat-api:3000 -t 0

exec "$@"