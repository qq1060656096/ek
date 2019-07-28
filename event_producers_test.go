package ek

import (
	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/assert"
	"testing"
)

func bTestNewEventProducers(t *testing.T) {

	//producers := NewProducers()
	//producers.SetAll(testProducersList)
	//clusters := NewClusters()
	//clusters.SetAll(testClustersList)
	//
	//ep := NewEventProducers(*producers, *clusters)
	//e := NewEventRaw("UserRegistr", "", make(map[string]string), ip.IpAddr(), time.Now().Unix())
	//producerName := "p0_aliyun_qywxSync"
	//syncProducer,err := ep.GetNewSyncProducer(producerName)
	//if err != nil {
	//
	//}
	//ep.SendSyncEvent(producerName, syncProducer, *e)
}

func GetTestConfigProducersClusters() (* Producers, *Clusters) {
	producerList := []*Producer{
		{
			Name: "c0DockerKafkaCluster_p0DockerKafkaProducer",
			ClusterName: "c0DockerKafkaCluster",
			TopicName: "test",
			Config: func() *sarama.Config{
				config := sarama.NewConfig()
				config.Producer.Return.Successes = true
				return config
			}(),
		},
	}
	producers := NewProducers()
	producers.SetAll(producerList)
	clusterList := []*Cluster{
		{
			Name: "c0DockerKafkaCluster",
			Addrs: []string{
				"199.199.199.199:9092",
			},
		},
	}
	clusters := NewClusters()
	clusters.SetAll(clusterList)
	return producers ,clusters
}

func TestNewEventProducers(t *testing.T) {
	producers, clusters := GetTestConfigProducersClusters()
	eventProducers := NewEventProducers(*producers, *clusters)
	_ , err := eventProducers.GetNewSyncProducer("c0DockerKafkaCluster_p0DockerKafkaProducer")
	assert.Equal(t, nil, err)
}

func TestEventProducers_SendSyncEvent(t *testing.T) {
	producers, clusters := GetTestConfigProducersClusters()
	eventProducers := NewEventProducers(*producers, *clusters)
	syncProducer , err := eventProducers.GetNewSyncProducer("c0DockerKafkaCluster_p0DockerKafkaProducer")
	assert.Equal(t, nil, err)
	event := NewEventRaw("UserRegister", "", map[string]interface{}{"a": 1}, "", 0)
	_, _, err = eventProducers.SendSyncEvent("c0DockerKafkaCluster_p0DockerKafkaProducer", syncProducer, event)
	assert.Equal(t, nil, err)
}

func TestEventProducers_SendSyncEvents(t *testing.T) {
	producers, clusters := GetTestConfigProducersClusters()
	eventProducers := NewEventProducers(*producers, *clusters)
	syncProducer , err := eventProducers.GetNewSyncProducer("c0DockerKafkaCluster_p0DockerKafkaProducer")
	assert.Equal(t, nil, err)
	var events []*Event
	events = make([]*Event, 2)
	events[0] = NewEventRaw("UserRegisterBatchEvent", "", map[string]interface{}{"a": 1}, "", 0)
	events[1] = NewEventRaw("UserRegisterBatchEvent", "", map[string]interface{}{"a": 1}, "", 0)
	err = eventProducers.SendSyncEvents("c0DockerKafkaCluster_p0DockerKafkaProducer", syncProducer, events)
	assert.Equal(t, nil, err)
}