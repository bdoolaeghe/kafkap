package main

/**
 * Copyright 2016 Confluent Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import (
	"fmt"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jawher/mow.cli"
	"os"
)

func main() {
	app := cli.App("kafkap", "Kafka producer")
	app.Spec = "[-b] -t -m [-h...]"
	var (
		broker  = app.StringOpt("b broker", "localhost:9092", "broker")
		topic   = app.StringOpt("t topic", "", "topic")
		body    = app.StringOpt("m body", "", "message body")
		headers = app.StringsOpt("h header", nil, "message header <key=value>")
	)

	app.Action = func() {
		fmt.Printf("\nSending message: '%v'", *body)
		fmt.Printf("\n   with headers: %v\n", *headers)

		producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": *broker})

		if err != nil {
			fmt.Printf("Failed to create producer: %s\n", err)
			os.Exit(1)
		}

		// Optional delivery channel, if not specified the Producer object's
		// .Events channel is used.
		deliveryChan := make(chan kafka.Event)

		// headers
		var kafkaHeaders []kafka.Header
		for _, element := range *headers {
			headerKeyValue := strings.Split(element, "=")
			headerKey := headerKeyValue[0]
			headerValue := headerKeyValue[1]
			newHeader := kafka.Header{Key: headerKey, Value: []byte(headerValue)}
			kafkaHeaders = append(kafkaHeaders, newHeader)
		}

		// produce
		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: topic, Partition: kafka.PartitionAny},
			Value:          []byte(*body),
			Headers:        kafkaHeaders,
		}, deliveryChan)

		e := <-deliveryChan
		m := e.(*kafka.Message)

		if m.TopicPartition.Error != nil {
			fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		} else {
			fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}

		close(deliveryChan)

	}

	app.Run(os.Args)
}
