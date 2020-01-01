FROM golang:1.13.5
WORKDIR /go/res-sys
COPY . .
RUN mkdir -p build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -extldflags "-static"' -o build/res-sys github.com/tmtx/res-sys/cmd/all/...

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/res-sys/build/res-sys .
COPY --from=0 /go/res-sys/.env .
CMD ["/bin/sh", "-c", "./res-sys -env production"]