FROM golang:alpine AS build

LABEL MAINTAINER = 'Friends of Go (it@friendsofgo.tech)'

RUN apk add --update git
WORKDIR /go/src/github.com/friendsofgo/workshop-microservices
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/counters-api cmd/counters-api/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/counters-api /go/bin/counters-api
ENTRYPOINT ["/go/bin/counters-api"]