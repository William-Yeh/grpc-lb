#======================================#
# builder env
#

FROM golang:1.14.1 AS builder
LABEL author="william.pjyeh@gmail.com"

WORKDIR /go/src/app
COPY .  .
RUN GOOS=linux   GOARCH=amd64  CGO_ENABLED=0 \
    go build -o server  server.go


#======================================#
# deployment env
#

FROM alpine:3.11.3
LABEL author="william.pjyeh@gmail.com"

# copy executable
COPY --from=builder /go/src/app/server  /

EXPOSE 80
EXPOSE 30051

# Define default command.
ENTRYPOINT ["/server"]
