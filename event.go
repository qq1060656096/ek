package ek

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"math/rand"
	"os"
	"time"
)

const (
	EventVersionDefault = "0"
)
type EventAdditional[]interface{}

type Event struct {
	Id string `json:"id"`
	Name string `json:"eventKey"`
	Key string `json:"key"`
	Version string `json:"v"`
	Ip string `json:"ip"`
	Time int64 `json:"time"`
	Data interface{} `json:"data"`
	Additional []EventAdditional `json:"addit"`
	// 转发到某个生产者
	ForwardProducerName string `json:"forward"`
	// 广播到某个生产者
	BroadcastProducerName string `json:"broadcast"`
	// 从那个生产者来
	ProducerName string `json:"producer"`
}

func NewEventRaw(name string, key string, data interface{}, ip string, opTime int64) *Event {
	return &Event{
		Id: GenerateEventId(ip),
		Name: name,
		Key: key,
		Version: EventVersionDefault,
		Ip: ip,
		Time: opTime,
		Data: data,
	}
}

func (e *Event) String() string {
	bytes, _ := json.Marshal(*e)
	return string(bytes)
}

func (e *Event) ToProducerMessage() *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic: e.Name,
		Value: sarama.StringEncoder(e.String()),
	}
	if e.Key != "" {
		msg.Key = sarama.StringEncoder(e.Key)
	}
	return msg
}

// AppendAdditional 增加附加信息
func (e *Event) AppendAdditional(additional EventAdditional) {
	e.Additional = append(e.Additional, additional)
}


// GenerateEventId 生成EventId
// 格式: 时间戳.纳秒.线程id.随机数.ip地址
func GenerateEventId(ip string) string {
	t := time.Now()
	return fmt.Sprintf("%d-%d-%d-%d-%s", t.Unix(), t.Nanosecond(), os.Getpid(), rand.Uint32(), ip)
}

