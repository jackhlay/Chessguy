FROM golang:1.22.10-bullseye

WORKDIR /app

COPY go.mod go.sum ./
COPY *.go ./

RUN go mod tidy && go mod download
RUN go build -o randoGames

CMD ["./randoGames"]

EXPOSE 5000
