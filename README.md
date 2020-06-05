# go-and-rabbitmq-simple-example

### Install RabbitMQ or use docker for get one instance
```
docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```

#### Run consumer to wait for messages, do not finish execution!
```
go run consumer.go
```

#### Run producer on other terminal for create a nice message
```
go run consumer.go
```

Paulo Ariel Pareja