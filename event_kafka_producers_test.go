package ek

import (
	"testing"
	"time"
	"github.com/qq1060656096/go-develop/pkg/ip"
)

func TestNewEventKafkaProducers(t *testing.T) {

	producers := NewProducers()
	producers.SetAll(testProducersList)
	clusters := NewClusters()
	clusters.SetAll(testClustersList)

	ekp := NewEventKafkaProducers(*producers, *clusters)
	e := NewEventRaw("UserRegistr", "", make(map[string]string), ip.IpAddr(), time.Now().Unix())
	producerName := "p0_aliyun_qywxSync"
	syncProducer,err := ekp.GetNewSyncProducer(producerName)
	if err != nil {

	}
	ekp.SendSyncEvent(producerName, syncProducer, *e)
}