package ek

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testClusters = map[string]Cluster{
	"c0_QywxApp":{
		Name: "c0_QywxApp",
		Addrs: []string{
			"199.199.199.199:9092",
			"127.0.0.1:9092",
		},
	},
	"c0_AliYun1":{
		Name: "c0_AliYun1",
		Addrs: []string{
			"199.199.199.199:9092",
			"127.0.0.1:9092",
		},
	},
	"c0_AliYun2":{
		Name: "c0_AliYun2",
		Addrs: []string{
			"199.199.199.199:9092",
			"127.0.0.1:9092",
		},
	},
}

var testClustersList = []*Cluster{
	{
		Name: "c0_QywxApp",
		Addrs: []string{
			"199.199.199.199:9092",
			"127.0.0.1:9092",
		},
	},
	{
		Name: "c0_AliYun1",
		Addrs: []string{
			"199.199.199.199:9092",
			"127.0.0.1:9092",
		},
	},
	{
		Name: "c0_AliYun2",
		Addrs: []string{
			"199.199.199.199:9092",
			"127.0.0.1:9092",
		},
	},
}
func TestClusters_GetAll(t *testing.T) {
	clusters := NewClusters()
	assert.Equal(t, 0, len(clusters.GetAll()))
}

func TestClusters_Set(t *testing.T) {
	clusters := NewClusters()
	for _, v := range testClusters {
		c := v
		clusters.Set(&c)
	}
	assert.Equal(t, 3, len(clusters.GetAll()))
}

func TestClusters_Get(t *testing.T) {
	clusters := NewClusters()
	c := testClusters["c0_QywxApp"]
	clusters.Set(&c)
	cNew, ok := clusters.Get(c.Name)
	assert.Equal(t, true, ok)
	assert.Equal(t, c.Name, cNew.Name)

	var cNil *Cluster
	cNew, ok = clusters.Get("not found cluster")
	assert.Equal(t, false, ok)
	assert.Equal(t, cNil, cNew)
}

func TestClusters_SetAll(t *testing.T) {
	clustersList := []*Cluster{}
	for _, v := range testClusters {
		c := v
		clustersList = append(clustersList, &c)
	}
	clusters := NewClusters()
	clusters.SetAll(clustersList)
	assert.Equal(t, len(clustersList), len(clusters.GetAll()))
}