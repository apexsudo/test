FROM golang:1.24 AS compiler
WORKDIR /app

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cli ./main.go

FROM scratch
COPY --from=compiler --chmod=111 /app/cli .
COPY --from=compiler /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./cli"]
