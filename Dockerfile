# golang alpine 1.13.5-alpine
FROM golang:1.16.2-alpine AS builder
# Create sngular user.
RUN adduser -D -g '' sngular
# Create workspace
WORKDIR /opt/app/
COPY go.mod .
# fetch dependancies
RUN go mod download
RUN go mod verify
# copy the source code as the last step
COPY . .
# build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/gitops-echobot .

# build a small image
FROM alpine:3.13.2
LABEL language="golang"
LABEL org.opencontainers.image.source https://github.com/sngular/gitops-echobot
# import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
# copy the static executable
COPY --from=builder /go/bin/gitops-echobot /usr/local/bin/gitops-echobot
# use an unprivileged user.
USER sngular
# run app
ENTRYPOINT ["gitops-echobot"]