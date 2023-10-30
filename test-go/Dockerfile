FROM golang:alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o /server

FROM alpine

COPY --from=builder /server /server

EXPOSE 8080

CMD ["/server"]
