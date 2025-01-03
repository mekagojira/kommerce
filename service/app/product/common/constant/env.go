package constant

import (
	"komo/lib/engine"
	"komo/lib/streaming"
)

func GetKafkaEnv() streaming.KafkaConfig {
	return streaming.KafkaConfig{ConsumerGroup: engine.GetEnv("PRODUCT_CONSUMER_GROUP", "CONSUMER_GROUP")}
}

func GetPgEnv() string {
	return engine.GetEnv("DATABASE_URI", "PRODUCT_DATABASE_URI")
}

func GetPortEnv() string {
	return engine.GetEnv("PRODUCT_PORT", "PORT", "9000")
}
