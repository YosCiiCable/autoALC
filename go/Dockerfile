FROM golang:1.20

# go run on docker
	WORKDIR /usr/src/app

	COPY go.mod go.sum ./
	RUN go mod download && go mod verify

	COPY . .
	RUN go build -v -o /usr/local/bin/app ./...


	RUN apt update
	RUN apt install ./aptFiles/google-chrome-stable_current_amd64.deb -y
	#RUN apt-get install google-chrome-stable -y

	CMD ["app"]
