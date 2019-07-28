# ek(Event Kafka)

## 生产者发送消息
**生产者配置**
```go
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
```
**生产者发送消息**
```go
// 请看测试文件: event_producers_test.go
eventProducers := NewEventProducers(*producers, *clusters)
var events []*Event
events = make([]*Event, 2)
events[0] = NewEventRaw("UserRegisterBatchEvent", "", map[string]interface{}{"a": 1}, "", 0)
events[1] = NewEventRaw("UserRegisterBatchEvent", "", map[string]interface{}{"a": 1}, "", 0)
// 批量发送消息
err = eventProducers.SendSyncEvents("c0DockerKafkaCluster_p0DockerKafkaProducer", syncProducer, events)
// 发送单条消息
event := NewEventRaw("UserRegister", "", map[string]interface{}{"a": 1}, "", 0)
_, _, err = eventProducers.SendSyncEvent("c0DockerKafkaCluster_p0DockerKafkaProducer", syncProducer, event)

```

```
# 单元测试
GO111MODULE=on GOPROXY=https://goproxy.io go test -v
```