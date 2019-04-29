# Example of Kafka logger in Go

The library [github.com/tracer0tong/kafkalogrus](github.com/tracer0tong/kafkalogrus) provides solution for forwarding logs to Kafka. The library implements custom [Logrus](https://github.com/sirupsen/logrus) hook which can be used in your project if your application uses Logrus logger.

Example:

```go
hook, err := kafkalogrus.NewKafkaLogrusHook(
    topic,
    logrus.AllLevels,
    &logrus.JSONFormatter{},
    brokers,
    topic,
    false,
    tlsConfig)
if err != nil {
    panic(err)
}
// Create a new logrus.Logger
logger := logrus.New()

// Add hook to logger
logger.Hooks.Add(hook)

// Send message to logger
logger.WithField("Field", "Value").Info("This is an Info msg")
logger.Warn("This is a Warn msg")
logger.Error("This is an Error msg")

```

There is example app which shows the setup:

```shell
➜  export KAFKA_BROKER="<broker>:<port>"
➜  export KAFKA_TOPIC="<topic>"      
➜  export KAFKA_CACERT="<path-to-ca-cert>"

➜  go run example_setup.go
INFO[0000] This is an Info msg                           Field=Value
WARN[0000] This is a Warn msg                           
ERRO[0000] This is an Error msg 


```