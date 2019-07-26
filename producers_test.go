package ek

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testProducers = map[string]Producer{
	"p0_AliYun1_QywxSync":{
		Name: "p0_AliYun1_QywxSync",
		ClusterName: "c0_AliYun1",
		TopicName: "qywx-sync",
	},
	"p0_AliYun1_User":{
		Name: "p0_AliYun1_User",
		ClusterName: "c0_AliYun1_user",
		TopicName: "user",

	},
}

var testProducersList = []*Producer{
	{
		Name: "p0_AliYun1_QywxSync",
		ClusterName: "c0_AliYun1",
		TopicName: "qywx-sync",
	},
	{
		Name: "p0_AliYun1_User",
		ClusterName: "c0_AliYun1_user",
		TopicName: "user",
	},
}

func TestProducers_GetAll(t *testing.T) {
	producers := NewProducers()
	assert.Equal(t, 0, len(producers.GetAll()))

	for _,v := range testProducers {
		tmp := v
		producers.Set(&tmp)
	}
}

func TestProducers_Set(t *testing.T) {
	producers := NewProducers()
	for _,v := range testProducers {
		tmp := v
		producers.Set(&tmp)
	}
	assert.LessOrEqual(t, 1, len(producers.GetAll()))
}

func TestProducers_Get(t *testing.T) {
	producers := NewProducers()
	var tmp Producer
	for _,v := range testProducers {
		tmp = v
		producers.Set(&tmp)
	}
	p, ok := producers.Get(tmp.Name)
	assert.Equal(t, true, ok)
	assert.Equal(t, tmp.Name, p.Name)
	assert.Equal(t, tmp.ClusterName, p.ClusterName)
}

func TestProducers_SetAll(t *testing.T) {
	producers := NewProducers()
	var tmp Producer
	var producersList []*Producer
	for _,v := range testProducers {
		tmp = v
		producersList = append(producersList, &tmp)
	}
	producers.SetAll(producersList)
	p, ok := producers.GetAll()[tmp.Name]
	assert.Equal(t, true, ok)
	assert.Equal(t, tmp.Name, p.Name)
	assert.Equal(t, tmp.ClusterName, p.ClusterName)
}