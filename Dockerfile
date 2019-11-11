FROM golang:alpine AS builder

# Install git for fetching dependencies
RUN apk update && apk add --no-cache git

WORKDIR /home_manager_heartbeat

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the binary.
RUN go build -o /go/bin/home_manager_heartbeat cmd/home_manager_heartbeat/main.go

## Build lighter image
FROM alpine:latest

# Copy our static executable.
COPY --from=builder /go/bin/home_manager_heartbeat /home_manager_heartbeat

# Run the binary.
ENTRYPOINT /home_manager_heartbeat