package ek

import (
	"github.com/Shopify/sarama"
)


// 事件生产者
type EventKafkaProducers struct {
	producers Producers
	clusters Clusters
}

// NewEventKafkaProducers 创建kafka事件生产者
func NewEventKafkaProducers(producers Producers, clusters Clusters) *EventKafkaProducers {
	return &EventKafkaProducers{
		producers: producers,
		clusters: clusters,
	}
}

//  GetClient 创建kafka client
func (ekp *EventKafkaProducers) GetClient(producerName string) (*sarama.Client, error) {
	producer, ok := ekp.producers.Get(producerName)
	if !ok{
		return nil, ErrProducerNotExist
	}
	cluster, ok := ekp.clusters.Get(producer.ClusterName)
	if !ok {
		return nil, ErrClusterNotExist
	}
	if producer.Config == nil {
		client, err := sarama.NewClient(cluster.Addrs, producer.Config)
		if err != nil {
			return nil, err
		}
		producer.Client = &client
	}
	return producer.Client, nil
}

// GetNewSyncProducer 创建同步kafka生产者
func (ekp *EventKafkaProducers) GetNewSyncProducer(producerName string) (*sarama.SyncProducer, error) {
	client, err := ekp.GetClient(producerName)
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
func (ekp *EventKafkaProducers) GetNewAsyncProducer(producerName string) (*sarama.SyncProducer, error) {
	client, err := ekp.GetClient(producerName)
	if err != nil {
		return nil, err
	}
	producer, err := sarama.NewSyncProducerFromClient(*client)
	if err != nil {
		return nil, err
	}
	return &producer, nil
}




// SendSyncEvent 发送同步event
func (ekp *EventKafkaProducers) SendSyncEvent (producerName string, syncProducer *sarama.SyncProducer, event Event) (partition int32, offset int64, err error) {
	producer, ok := ekp.producers.Get(producerName)
	if !ok{
		err = ErrProducerNotExist
		return
	}
	event.ProducerName = producer.Name
	return (*syncProducer).SendMessage(event.ToProducerMessage())
}

// SendSyncEvents 发送同步events
func (ekp *EventKafkaProducers) SendSyncEvents (producerName string, syncProducer *sarama.SyncProducer, events []Event) (err error) {
	producer, ok := ekp.producers.Get(producerName)
	if !ok{
		err = ErrProducerNotExist
		return
	}
	var msgs []*sarama.ProducerMessage
	for _, event := range events {
		event.ProducerName = producer.Name
		msgs = append(msgs, event.ToProducerMessage())
	}
	return (*syncProducer).SendMessages(msgs)
}


// SendAsyncEvent 发送异步event
func (ekp *EventKafkaProducers) SendAsyncEvent(producerName string, asyncProducer *sarama.AsyncProducer, event Event) (err error) {
	producer, ok := ekp.producers.Get(producerName)
	if !ok{
		err = ErrProducerNotExist
		return
	}
	event.ProducerName = producer.Name
	(*asyncProducer).Input() <- event.ToProducerMessage()
	return nil
}

// SendAsyncEvents 发送异步events
func (ekp *EventKafkaProducers) SendAsyncEvents(producerName string, asyncProducer *sarama.AsyncProducer, events []Event) (err error) {
	producer, ok := ekp.producers.Get(producerName)
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