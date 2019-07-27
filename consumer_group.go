package ek

import "github.com/Shopify/sarama"

type ConsumerGroup struct {
	Name string
	GroupName string
	TopicsName []string
	ClusterName string
	Client *sarama.Client
	Config *sarama.Config
	EventConsumerGroupHandler *EventConsumerGroupHandler
	// 转发到某个生产者
	ForwardProducerName string `json:"forward"`
	// 广播到来源生产者
	BroadcastOriginProducerName string `json:"broadcast"`
	// 广播到指定生产者
	BroadcastProducerName string `json:"producer"`
}
