package streaming

import (
	"context"
	"fmt"
	"komo/lib/engine"
	"strings"
	"sync"

	"github.com/twmb/franz-go/pkg/kgo"
)

var Kafka *kgo.Client

type KafkaConfig struct {
	ConsumerGroup string
}

func Connect(config KafkaConfig) {
	brokers := engine.GetEnv("KAFKA_BROKER_URI", "localhost:9092")
	seeds := strings.Split(brokers, ",")

	client, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
		kgo.ConsumerGroup(config.ConsumerGroup),
	)

	if err != nil {
		panic(err)
	}

	Kafka = client
	if err := Kafka.Ping(context.Background()); err != nil {
		panic(err)
	}

	engine.Logger.Info("Kafka connected")
}

func Test() {
	ctx := context.Background()

	var wg sync.WaitGroup
	wg.Add(1)
	record := &kgo.Record{Topic: "komo_product", Value: []byte("bar")}
	Kafka.Produce(ctx, record, func(record *kgo.Record, err error) {
		defer wg.Done()
		if err != nil {
			fmt.Printf("record had a produce error: %v\n", err)
		}

		fmt.Println(record)

	})
	wg.Wait()
}
