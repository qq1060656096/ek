package ek

import "github.com/Shopify/sarama"

// 消费组默认配置
func NewConsumerGroupDefaultConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V1_0_0_0
	config.Consumer.Return.Errors = true
	return config
}