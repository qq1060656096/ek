package ek

import (
	"github.com/Shopify/sarama"
)

func init()  {

}

// Default 默认 事件kafka生产者
var Default *EventProducers

// 事件生产者
type EventProducers struct {
	producers Producers
	clusters Clusters
}



// NewEventProducers 创建kafka事件生产者
func NewEventProducers(producers Producers, clusters Clusters) *EventProducers {
	return &EventProducers{
		producers: producers,
		clusters: clusters,
	}
}

//  GetClient 创建kafka client
func (ep *EventProducers) GetClient(producerName string) (*sarama.Client, error) {
	producer, ok := ep.producers.Get(producerName)
	if !ok{
		return nil, ErrProducerNotExist
	}
	cluster, ok := ep.clusters.Get(producer.ClusterName)
	if !ok {
		return nil, ErrClusterNotExist
	}
	if producer.Client == nil {
		client, err := sarama.NewClient(cluster.Addrs, producer.Config)
		if err != nil {
			return nil, err
		}
		producer.Client = &client
	}
	return producer.Client, nil
}

// GetNewSyncProducer 创建同步kafka生产者
func (ep *EventProducers) GetNewSyncProducer(producerName string) (*sarama.SyncProducer, error) {
	client, err := ep.GetClient(producerName)
	if err != nil {
		return nil, err
	}
	producer, err := sarama.NewSyncProducerFromClient(*client)
	if err != nil {
		return nil, err
	}
	return &producer, nil
}

// GetNewAsyncProducer 创建异步kafka生产者
func (ep *EventProducers) GetNewAsyncProducer(producerName string) (*sarama.AsyncProducer, error) {
	client, err := ep.GetClient(producerName)
	if err != nil {
		return nil, err
	}
	producer, err := sarama.NewAsyncProducerFromClient(*client)
	if err != nil {
		return nil, err
	}
	return &producer, nil
}




// SendSyncEvent 发送同步event
func (ep *EventProducers) SendSyncEvent (producerName string, syncProducer *sarama.SyncProducer, event *Event) (partition int32, offset int64, err error) {
	producer, ok := ep.producers.Get(producerName)
	if !ok{
		err = ErrProducerNotExist
		return
	}
	event.ProducerName = producer.Name
	msg := event.ToProducerMessage()
	msg.Topic = producer.TopicName
	return (*syncProducer).SendMessage(msg)
}

// SendSyncEvents 发送同步events
func (ep *EventProducers) SendSyncEvents (producerName string, syncProducer *sarama.SyncProducer, events []*Event) (err error) {
	producer, ok := ep.producers.Get(producerName)
	if !ok{
		err = ErrProducerNotExist
		return
	}
	var msgs []*sarama.ProducerMessage
	for _, event := range events {
		event.ProducerName = producer.Name
		msg := event.ToProducerMessage()
		msg.Topic = producer.TopicName
		msgs = append(msgs, msg)
	}
	return (*syncProducer).SendMessages(msgs)
}


// SendAsyncEvent 发送异步event
func (ep *EventProducers) SendAsyncEvent(producerName string, asyncProducer *sarama.AsyncProducer, event *Event) (err error) {
	producer, ok := ep.producers.Get(producerName)
	if !ok{
		err = ErrProducerNotExist
		return
	}
	event.ProducerName = producer.Name
	msg := event.ToProducerMessage()
	msg.Topic = producer.TopicName
	(*asyncProducer).Input() <- msg
	return nil
}

// SendAsyncEvents 发送异步events
func (ep *EventProducers) SendAsyncEvents(producerName string, asyncProducer *sarama.AsyncProducer, events []*Event) (err error) {
	producer, ok := ep.producers.Get(producerName)
	if !ok{
		err = ErrProducerNotExist
		return
	}
	for _, event := range events {
		msg := &sarama.ProducerMessage{
			Topic: producer.TopicName,
			Value: sarama.StringEncoder(event.String()),
		}
		if event.Key != "" {
			msg.Key = sarama.StringEncoder(event.Key)
		}
		(*asyncProducer).Input() <-msg
	}
	return
}