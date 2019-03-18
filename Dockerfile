FROM golang:alpine as builder
RUN apk update && apk add git
ADD . /build/
WORKDIR /build/cmd/server
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
COPY --from=builder /build/cmd/server/main /app/
WORKDIR /app
CMD ["./main"]