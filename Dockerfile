FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

EXPOSE 8081

COPY . .

RUN go mod tidy
RUN go build -o binary cmd/main.go

ENTRYPOINT ["/app/binary"]
