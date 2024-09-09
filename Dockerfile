FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /gogo

# Final minimal image
FROM alpine:3.18
COPY --from=builder /gogo /gogo
EXPOSE 8080
CMD ["/gogo"]
