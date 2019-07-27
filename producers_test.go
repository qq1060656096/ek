package ek

import (
	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/assert"
	"testing"
)


var testProducers = []*Producer{
	{
		Name: "cAliYun_pBbs",
		ClusterName: "cAliYun",
		TopicName: "tBbs",
		Config: &sarama.Config{},
	},
	{
		Name: "cAliYun_pUser",
		ClusterName: "cAliYun",
		TopicName: "tUser",
		Config: &sarama.Config{},
	},
}

func TestProducers_GetAll(t *testing.T) {
	producers := NewProducers()
	assert.Equal(t, 0, len(producers.GetAll()))
}

func TestProducers_Set(t *testing.T) {
	producers := NewProducers()
	for _,v := range testProducers {
		producers.Set(v)
	}
	assert.Equal(t, len(testProducers), len(producers.GetAll()))
}

func TestProducers_Get(t *testing.T) {
	producers := NewProducers()
	var tmp *Producer
	for _,v := range testProducers {
		tmp = v
		producers.Set(tmp)
	}
	p, ok := producers.Get(tmp.Name)
	assert.Equal(t, true, ok)
	assert.Equal(t, tmp.Name, p.Name)
	assert.Equal(t, tmp.ClusterName, p.ClusterName)
}

func TestProducers_SetAll(t *testing.T) {
	producers := NewProducers()
	tmp := testProducers[0]
	producers.SetAll(testProducers)
	p, ok := producers.GetAll()[tmp.Name]
	assert.Equal(t, true, ok)
	assert.Equal(t, tmp.Name, p.Name)
	assert.Equal(t, tmp.ClusterName, p.ClusterName)
}