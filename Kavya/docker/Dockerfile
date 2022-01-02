FROM golang:alpine3.13 

RUN mkdir /app
WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o /bin/readwrite .

FROM alpine:latest
WORKDIR /
COPY --from=0 /bin/readwrite /readwrite
CMD [ "/readwrite" ]