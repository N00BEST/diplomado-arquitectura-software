FROM golang:1.27 as builder

WORKDIR /app

COPY go.mod go.sum ./
COPY main.go db.go records.go ./
COPY migrations/ ./migrations/

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o /actividad-03 .

FROM alpine:latest

RUN adduser -D nonroot
USER nonroot

WORKDIR /app

COPY --from=builder /actividad-03 /actividad-03

ENV GIN_MODE=release

EXPOSE 3000

ENTRYPOINT ["/actividad-03"]