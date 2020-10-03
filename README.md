# Microservices Demo for Workshop

This repository has the purpose of acting like a guide through our workshop about Microservices in Go.

This repository is split in many branches, each branch represents a final part of each iteration.

## Counters API

The application we will create during the workshop is a microservices, part of a bigger project about counters. Main goal of this workshop is to offer a demostration of how you can create a simple but well organized microservice using Go.

It is built using [Gorilla mux](https://github.com/gorilla/mux), a powerful library to create APIs.

The final result will be a fully functional microservice, that will use [mongoDB](https://www.mongodb.com/es)
for storing the data, and [Kafka](https://kafka.apache.org/) for sharing events between our microservices and the others.

## Libraries

In this application we're using, of course, the standard library, but also some third party libraries as:

* For mongoDB: [go.mongodb.org/mongo-driver](https://github.com/mongodb/mongo-go-driver)
* For Kafka: [github.com/segmentio/kafka-go](https://github.com/segmentio/kafka-go)
* For Logging: [github.com/uber-go/zap](https://github.com/uber-go/zap)

## Using our microservices

### Run docker (only necessary for solution two onwards)
First of all you need a Lenses key: it's totally free, you only need to register on: [https://lenses.io/downloads/lenses/](https://lenses.io/downloads/lenses/)
After doing it you'll receive a message with a link, that will generate your key.

Once you have the key, you just need to replace it on the `docker-compose.yml`, `{LENSE_ID}`.

```sh
$ docker-compose up -d
```

### Build

```sh
$ make build
```

### Launch

```sh
$ make run
```

For testing the application we're using a tool that is using `go test` behind the scenes, but also puts colors on it.
For that you need to install:

```sh
GO111MODULE=off go get -u github.com/rakyll/gotest
```

And then you can run:

```
make test
```

## License
MIT License, see [LICENSE](https://github.com/friendsofgo/workshop-microservices/blob/master/LICENSE)


  
