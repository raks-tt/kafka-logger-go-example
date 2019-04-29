package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/tracer0tong/kafkalogrus"

	"github.com/sirupsen/logrus"
)

func loadTLSConfig(caCertPath string) *tls.Config {
	// Load CA cert
	caCert, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}
	return tlsConfig
}

func main() {
	brokers := strings.Split(os.Getenv("KAFKA_BROKER"), ";")
	caCert := os.Getenv("KAFKA_CACERT")
	topic := os.Getenv("KAFKA_TOPIC")

	tlsConfig := loadTLSConfig(caCert)

	// Create a new hook
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

	// Ensure log messages were written to Kafka
	time.Sleep(time.Second)
}
