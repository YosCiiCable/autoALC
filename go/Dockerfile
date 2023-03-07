FROM golang:1.20

# go run on docker
	WORKDIR /usr/src/app

	COPY go.mod go.sum ./
	RUN go mod download && go mod verify

	COPY . .
	RUN go build -v -o /usr/local/bin/app ./...


	RUN sh -c 'echo "deb http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list'
	RUN wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add -
	RUN apt update
	RUN apt-get install google-chrome-stable -y

	CMD ["app"]
