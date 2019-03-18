FROM golang:alpine as builder
RUN apk update && apk add git && apk add ca-certificates
ADD . /build/
WORKDIR /build/cmd/server
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
RUN apk add top
COPY --from=builder /build/cmd/server/main /app/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
CMD ["./main"]