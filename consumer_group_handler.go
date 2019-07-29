package ek

import (
	"fmt"
	"github.com/Shopify/sarama"
)

type EventConsumeFunc func (msg *sarama.ConsumerMessage, event *Event)

type EventConsumeResult struct {
	Status string
	Message string
}


type EventConsumerGroupHandler struct {
	sarama.ConsumerGroupHandler
	EventConsumeFunc map[string] EventConsumeFunc
}

// Setup 在新会话开始时运行,即在消费品目标之前
func (o EventConsumerGroupHandler) Setup(sess sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup 清理是在会话结束时运行的,一旦所有的consumereclaim goroutine都退出,但是在最后一次提交偏移量之前
func (o EventConsumerGroupHandler) Cleanup(sess sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim 消费事件
func (o EventConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		t := msg
		event, err := ConsumerMessageToEvent(t)
		if err != nil {

		}
		eventConsumeFunc, ok := o.EventConsumeFunc[event.Name]
		// 没有配置事件
		if !ok {
			fmt.Println("consumer_group event no config", event.Name)
			continue
		}
		// 消费事件
		eventConsumeFunc(t, event)

		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		// 标记消息已经消费
		sess.MarkMessage(msg, "")
	}
	return nil
}