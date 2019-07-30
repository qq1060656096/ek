package ek

import "github.com/Shopify/sarama"

// NewSyncProducerDefaultConfig 创建生产者默认配置
func NewSyncProducerDefaultConfig() *sarama.Config {
	config := NewAsyncProducerDefaultConfig()
	config.Producer.Return.Successes = true
	return config
}

// NewAsyncProducerDefaultConfig 异步生产者配置
func NewAsyncProducerDefaultConfig() *sarama.Config {
	return sarama.NewConfig()
}


// 消费组默认配置
func NewConsumerGroupDefaultConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V1_0_0_0
	config.Consumer.Return.Errors = true
	return config
}