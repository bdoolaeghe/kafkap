Golang Kafka CLI for sending message
====================================
**kafkap** is a Kafka CLI for sending message (body & headers) on a kafka topic.

Install
-------
* Make sure to dispose of [golang SDK](https://github.com/golang/go/wiki/Ubuntu)
* install  librdkafka v0.11.5 or later:
```
# Install the Confluent public key
wget -qO - https://packages.confluent.io/deb/5.0/archive.key | sudo apt-key add -
# Add the repository to your /etc/apt/sources.list 
sudo add-apt-repository "deb [arch=amd64] https://packages.confluent.io/deb/5.0 stable main"
# update apt cache and install librdkafka-dev
sudo apt-get update && sudo apt-get install librdkafka-dev
```
* clone repo and build kafkap
```
make setup
make build
```

Usage
-----
Run `kafkap` for usage. 

Links
-----
based on [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go) [examples](https://github.com/confluentinc/confluent-kafka-go/blob/master/examples/producer_example/producer_example.go).
