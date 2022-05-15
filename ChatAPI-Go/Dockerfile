# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download


RUN mkdir -p /go/src/github.com/samra10/chat-api-go

COPY . /go/src/github.com/samra10/chat-api-go

COPY entrypoint.sh /usr/bin/entrypoint-go.sh
RUN chmod +x /usr/bin/entrypoint-go.sh

COPY wait-for-it.sh /usr/bin/
RUN chmod +x /usr/bin/wait-for-it.sh

ENTRYPOINT [ "entrypoint-go.sh" ]

EXPOSE 8080

CMD ["/go/src/github.com/samra10/chat-api-go/cmd/chat-api-go/chat-api-go"]