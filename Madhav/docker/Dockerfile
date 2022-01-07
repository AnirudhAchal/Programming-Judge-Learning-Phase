FROM golang:1.16-alpine3.13

RUN mkdir /app
WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o /bin/diner .

ENTRYPOINT [ "diner" ]