package ek

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetTestConfigConsumerGroupsClusters() ( *ConsumerGroups, *Clusters) {
	consumerList := []*ConsumerGroup{
		{
			Name: "c0DockerKafkaCluster_g0DockerKafkaConsumerGroups",
			GroupName: "c0DockerKafkaCluster_g0DockerKafkaConsumerGroups",
			ClusterName: "c0DockerKafkaCluster",
			TopicsName: []string{"test"},
			Config: NewConsumerGroupDefaultConfig(),
			EventConsumerGroupHandler: EventConsumerGroupHandler{
				EventConsumeFunc: map[string]EventConsumeFunc{
					"UserRegisterSyncSend": func(msg *sarama.ConsumerMessage, event *Event) {
						fmt.Println(event.String())
					},
				},
			},
		},
	}
	consumers := NewConsumerGroups()
	consumers.SetAll(consumerList)
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
	return consumers ,clusters
}

func TestEventConsumerGroups_ConsumeEvent(t *testing.T) {
	consumers ,clusters := GetTestConfigConsumerGroupsClusters()
	consumerGroups := NewEventConsumerGroups(*consumers, *clusters)
	consumerGroup, err := consumerGroups.GetNewConsumerGroup("c0DockerKafkaCluster_g0DockerKafkaConsumerGroups")
	assert.Equal(t, nil, err)
	consumerGroups.ConsumeEvent("c0DockerKafkaCluster_g0DockerKafkaConsumerGroups", consumerGroup)
}