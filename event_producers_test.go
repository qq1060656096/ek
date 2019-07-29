package ek

import (
	"github.com/stretchr/testify/assert"
	"log"
	"runtime"
	"testing"
	"time"
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
			Config: NewSyncProducerDefaultConfig(),
		},
		{
			Name: "c0DockerKafkaCluster_p0DockerKafkaAsyncProducer",// 异步生产者
			ClusterName: "c0DockerKafkaCluster",
			TopicName: "test",
			Config: nil,
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
	ip := "127.0.0.1"
	var opTime int64 = 1564045436// 2019-07-25 17:03:56
	event := NewEventRaw("UserRegisterSyncSend", "", map[string]interface{}{"a": 1}, ip, opTime)
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
	events[0] = NewEvent("UserRegisterBatchEventSyncSend", "", map[string]interface{}{"a": 1})
	events[1] = NewEvent("UserRegisterBatchEventSyncSend", "", map[string]interface{}{"a": 2})
	err = eventProducers.SendSyncEvents("c0DockerKafkaCluster_p0DockerKafkaProducer", syncProducer, events)
	assert.Equal(t, nil, err)
}

func TestEventProducers_SendAsyncEvent(t *testing.T) {
	producers, clusters := GetTestConfigProducersClusters()
	eventProducers := NewEventProducers(*producers, *clusters)
	asyncProducer , err := eventProducers.GetNewAsyncProducer("c0DockerKafkaCluster_p0DockerKafkaAsyncProducer")
	assert.Equal(t, nil, err)
	event := NewEvent("UserRegisterEventAsyncSend", "", map[string]interface{}{"a": 1})
	err = eventProducers.SendAsyncEvent("c0DockerKafkaCluster_p0DockerKafkaAsyncProducer", asyncProducer, event)
	for {
		runtime.Gosched()
		select {
		case err := <-(*asyncProducer).Errors():
			log.Println("Failed to produce message", err)
			assert.Equal(t, nil, err)
		default:
		}
		time.Sleep(time.Millisecond * 100)
		runtime.Gosched()
		break
	}
	assert.Equal(t, nil, err)
}

func TestEventProducers_SendAsyncEvents(t *testing.T) {
	producers, clusters := GetTestConfigProducersClusters()
	eventProducers := NewEventProducers(*producers, *clusters)
	asyncProducer , err := eventProducers.GetNewAsyncProducer("c0DockerKafkaCluster_p0DockerKafkaAsyncProducer")
	assert.Equal(t, nil, err)
	var events = make([]*Event, 2)
	events[0] = NewEvent("UserRegisterBatchEventAsyncSend", "", map[string]interface{}{"async": 21})
	events[1] = NewEvent("UserRegisterBatchEventAsyncSend", "", map[string]interface{}{"async": 22})
	err = eventProducers.SendAsyncEvents("c0DockerKafkaCluster_p0DockerKafkaAsyncProducer", asyncProducer, events)
	for {
		runtime.Gosched()
		select {
		case err := <-(*asyncProducer).Errors():
			log.Println("Failed to produce message", err)
			assert.Equal(t, nil, err)
		default:
		}
		time.Sleep(time.Millisecond * 100)
		runtime.Gosched()
		break
	}
	assert.Equal(t, nil, err)
}