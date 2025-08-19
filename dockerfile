FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /name-of-project

EXPOSE 8080

CMD ["/CHANGE-TO-NAME-OF-PROJECT"]
