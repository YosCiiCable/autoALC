FROM golang:1.20

# go run on docker
	WORKDIR /usr/src/app

	COPY go/go.mod go/go.sum ./
	RUN go mod download && go mod verify

	COPY ./go/* .
	RUN go build -v -o /usr/local/bin/app ./...
