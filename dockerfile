# Step 1: Build stage
FROM golang:1.24.2 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o myapp ./cmd

# Step 2: Run stage
FROM debian:bookworm-slim

# Install netcat for wait-for.sh
RUN apt-get update && apt-get install -y netcat-openbsd && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=build /app/myapp .

RUN chmod +x myapp

CMD ["./myapp"]
