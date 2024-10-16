FROM golang:1.22.2-alpine3.19 as build

RUN apk add --update --no-cache git

RUN mkdir /app
WORKDIR /app

COPY go.mod ./

RUN go mod tidy

COPY . .

RUN go build -o go-todo-api main.go

FROM alpine:3.19
WORKDIR /app
COPY --from=build /app/go-todo-api /app/go-todo-api 

CMD ["/app/go-todo-api"]