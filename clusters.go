package ek

// Clusters kafka集群
type Clusters struct {
	m map[string]*Cluster
}

func NewClusters() *Clusters{
	return &Clusters{
		m: make(map[string]*Cluster),
	}
}
func (c *Clusters) GetAll() map[string]*Cluster {
	return c.m
}

func (c *Clusters) Get(name string) (*Cluster, bool) {
	v, ok := c.m[name]
	return v, ok
}

func (c *Clusters) SetAll(clusters []*Cluster)  {
	for _, v := range clusters {
		c.m[v.Name] = v
	}
}

func (c *Clusters) Set(cluster *Cluster) {
	c.m[cluster.Name] = cluster
}