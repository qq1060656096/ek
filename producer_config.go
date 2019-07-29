package ek

import "github.com/Shopify/sarama"

// NewSyncProducerDefaultConfig 创建生产者默认配置
func NewSyncProducerDefaultConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	return config
}