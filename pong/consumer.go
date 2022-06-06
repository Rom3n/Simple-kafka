package main

import (
	kafka "github.com/segmentio/kafka-go"
	"log"
)

func main() {
	/*
		cfg, err := config.FromEnv()
		if err != nil {
			panic(err)
		}

		cfg.ServiceName = "ping"

		tracer, closer, err := cfg.NewTracer()
		if err != nil {
			panic(err)
		}

		defer closer.Close()

		opentracing.SetGlobalTracer(tracer)
	*/
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9093"),
		Topic:    "test-topic",
		Balancer: &kafka.LeastBytes{},
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

}
