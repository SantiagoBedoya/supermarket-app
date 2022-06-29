FROM golang:1.18-alpine AS builder
WORKDIR /app

COPY . .

RUN go build -o main main.go

FROM alpine

WORKDIR /app
COPY --from=builder /app/main .
ARG PORT
EXPOSE ${PORT}
CMD ["./main"]