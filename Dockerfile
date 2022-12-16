FROM golang:alpine
LABEL maintainer="Mrudraia"

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

RUN mkdir /app
WORKDIR /app

COPY .env .
COPY go.mod go.sum ./

RUN go get -d -v ./...

RUN go install -v ./...

COPY . .
# RUN go build -o /build

EXPOSE 9090

CMD ["go", "run", "./podlist/main.go"]