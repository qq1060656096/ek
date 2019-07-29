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
events[0] = NewEvent("UserRegisterBatchEvent", "", map[string]interface{}{"a": 1})
events[1] = NewEvent("UserRegisterBatchEvent", "", map[string]interface{}{"a": 1})
// 批量发送消息
err = eventProducers.SendSyncEvents("c0DockerKafkaCluster_p0DockerKafkaProducer", syncProducer, events)
// 发送单条消息
event := NewEvent("UserRegister", "", map[string]interface{}{"a": 1})
_, _, err = eventProducers.SendSyncEvent("c0DockerKafkaCluster_p0DockerKafkaProducer", syncProducer, event)

```

**消费者配置**
```go
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
```

**消费者消费**
```go

```
// 创建消费组
consumerGroups := NewEventConsumerGroups(*consumers, *clusters)
// 创建单个消费者
consumerGroup, err := consumerGroups.GetNewConsumerGroup("c0DockerKafkaCluster_g0DockerKafkaConsumerGroups")
// 消费消息
consumerGroups.ConsumeEvent("c0DockerKafkaCluster_g0DockerKafkaConsumerGroups", consumerGroup)
```
# 单元测试
GO111MODULE=on GOPROXY=https://goproxy.io go test -v
```