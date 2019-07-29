package ek

import (
	"context"
	"github.com/Shopify/sarama"
)


// 事件消费者
type EventConsumerGroups struct {
	consumerGroups ConsumerGroups
	clusters Clusters
	eventConsumerGroupHandler *EventConsumerGroupHandler
}



// NewEventConsumerGroups 创建kafka事件消费者
func NewEventConsumerGroups(consumerGroups ConsumerGroups, clusters Clusters) *EventConsumerGroups {
	return &EventConsumerGroups{
		consumerGroups: consumerGroups,
		clusters: clusters,
	}
}

//  GetClient 创建kafka client
func (ecg *EventConsumerGroups) GetClient(consumerName string) (*sarama.Client, error) {
	consumer, ok := ecg.consumerGroups.Get(consumerName)
	if !ok{
		return nil, ErrConsumerGroupNotExist
	}
	cluster, ok := ecg.clusters.Get(consumer.ClusterName)
	if !ok {
		return nil, ErrClusterNotExist
	}
	if consumer.Client == nil {
		client, err := sarama.NewClient(cluster.Addrs, consumer.Config)
		if err != nil {
			return nil, err
		}
		consumer.Client = &client
	}
	return consumer.Client, nil
}

// GetNewConsumerGroup 创建kafka消费者
func (ecg *EventConsumerGroups) GetNewConsumerGroup(consumerName string) (*sarama.ConsumerGroup, error) {
	consumer, ok := ecg.consumerGroups.Get(consumerName)
	if !ok{
		return nil, ErrConsumerGroupNotExist
	}
	client, err := ecg.GetClient(consumerName)
	if err != nil {
		return nil, err
	}
	consumerGroup, err := sarama.NewConsumerGroupFromClient(consumer.GroupName, *client)
	if err != nil {
		return nil, err
	}
	return &consumerGroup, nil
}



// ConsumeEvent 消费者消费事件(注意会堵塞)
func (ecg *EventConsumerGroups) ConsumeEvent (consumerName string, consumerGroup *sarama.ConsumerGroup) error {
	consumer, ok := ecg.consumerGroups.Get(consumerName)
	if !ok{
		return ErrClusterNotExist
	}
	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		err := (*consumerGroup).Consume(ctx, consumer.TopicsName, consumer.EventConsumerGroupHandler)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

// MustConsumeEvent 消费者消费事件(注意会堵塞)
func (ecg *EventConsumerGroups) MustConsumeEvent (consumerName string, consumerGroup *sarama.ConsumerGroup, event Event) error {
	consumer, ok := ecg.consumerGroups.Get(consumerName)
	if !ok{
		return ErrClusterNotExist
	}
	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		err := (*consumerGroup).Consume(ctx, consumer.TopicsName, consumer.EventConsumerGroupHandler)
		if err != nil {
			panic(err)
		}
	}
	return nil
}