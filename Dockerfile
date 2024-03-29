FROM golang:latest as builder
LABEL maintainer="koteyye@yandex.ru"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o apitraning ./cmd/apitraning/main.go

######## Start a new stage from scratch #######
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/apitraning .
CMD ["./apitraning"]
