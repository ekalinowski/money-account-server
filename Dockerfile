# Build from golang latest image
FROM golang:latest as builder
WORKDIR /server
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# Build api image from scratch, copying pre-built binary and files
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /server/server ./

EXPOSE 8080
ENTRYPOINT ["./server"]