package ek

// Producers kafka集群
type ConsumerGroups struct {
	m map[string]*ConsumerGroup
}

func NewConsumerGroups() *ConsumerGroups{
	return &ConsumerGroups{
		m: make(map[string]*ConsumerGroup),
	}
}
func (c *ConsumerGroups) GetAll() map[string]*ConsumerGroup {
	return c.m
}

func (c *ConsumerGroups) Get(name string) (*ConsumerGroup, bool) {
	v, ok := c.m[name]
	return v, ok
}

func (c *ConsumerGroups) SetAll(ConsumerGroups []*ConsumerGroup)  {
	for _, v := range ConsumerGroups {
		c.m[v.Name] = v
	}
}

func (c *ConsumerGroups) Set(ConsumerGroup *ConsumerGroup) {
	c.m[ConsumerGroup.Name] = ConsumerGroup
}