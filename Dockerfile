FROM golang:alpine AS build_base

WORKDIR /tmp/go_app

COPY . .

RUN go mod download
RUN go build -o ./hangman_player .

FROM alpine:latest

COPY --from=build_base /tmp/go_app/hangman_player /app/hangman_player

EXPOSE 8090

CMD ["/app/hangman_player"]