package ek

import "github.com/Shopify/sarama"

type Producer struct {
	Name string
	ClusterName string
	TopicName string
	Client *sarama.Client
	Config *sarama.Config
}
