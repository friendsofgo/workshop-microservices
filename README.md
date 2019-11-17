# Microservices Demo for Workshop

This repository have the purpose of serve of guide through to our workshop about Microservices on Go.

This repository is divide by branches, each branch represent a final part of each sprint.

## Counters API

The application that will be created during the workshop is a one microservices of a more big project about
counters. With this workshop we can offer a demostration of how you can create a simple microservice, using
Go.

It is built using [Gorilla mux](https://github.com/gorilla/mux), a powerful library to create APIs.

The final microservices will be fully functional, for that we use [mongoDB](https://www.mongodb.com/es)
to store the data, and [Kafka](https://kafka.apache.org/) for sharing events between our microservices and the others.

## Libraries

In this application we're using, of course, the standard library, but also some third party libraries as:

* For mongoDB: [go.mongodb.org/mongo-driver](https://github.com/mongodb/mongo-go-driver)
* For Kafka: [github.com/segmentio/kafka-go](https://github.com/segmentio/kafka-go)
* For Logging: [github.com/uber-go/zap](https://github.com/uber-go/zap)

## Using our microservices

### Run docker (only necessary for solution two onwards)
First of all you need a lenses key, it's totally free you only need to register on: [https://lenses.io/downloads/lenses/](https://lenses.io/downloads/lenses/)
then they send you a message with a link, this link will have your necessary key.

You need to replace on the `docker-compose.yml`, `{LENSE_ID}` for your key and that's all.

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

For testing the application we're using a tool that using `go test` underlying but put colors on it.
For that you need to install before to run it:

```sh
GO111MODULE=off go get -u github.com/rakyll/gotest
```

And then you can run:

```
make test
```

## License
MIT License, see [LICENSE](https://github.com/friendsofgo/workshop-microservices/blob/master/LICENSE)


  
