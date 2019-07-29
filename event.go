package ek

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"math/rand"
	"os"
	"time"
	"github.com/qq1060656096/go-develop/pkg/ip"
)

const (
	EventVersionDefault = "0"
)
var localIp *ip.Ip

type EventAdditional[]interface{}

// Event 事件对象
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


func init() {
	var err error
	localIp, err = ip.New()
	if err != nil {
		log.Println("package.ek.event.init.error: ",err)
	}
}

func NewEvent(name string, key string, data interface{}) *Event {
	ipAddr := localIp.IpAddr()
	return NewEventRaw(name, key, data, ipAddr, time.Now().Unix())
}

// NewEventRaw 创建事件
func NewEventRaw(name string, key string, data interface{}, ipAddr string, opTime int64) *Event {
	return &Event{
		Id: GenerateEventId(ipAddr),
		Name: name,
		Key: key,
		Version: EventVersionDefault,
		Ip: ipAddr,
		Time: opTime,
		Data: data,
	}
}

// String 事件转换成string
func (e *Event) String() string {
	bytes, _ := json.Marshal(*e)
	return string(bytes)
}

// ToProducerMessage 事件转ProducerMessage对象
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

// ConsumerMessageToEvent ConsumerMessageToEvent 对象转Event事件
func ConsumerMessageToEvent(msg *sarama.ConsumerMessage) (*Event, error) {
	event := &Event{
	}
	err := json.Unmarshal(msg.Value, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}
