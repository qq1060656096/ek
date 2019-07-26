package ek

// Producers kafka集群
type Producers struct {
	m map[string]*Producer
}

func NewProducers() *Producers{
	return &Producers{
		m: make(map[string]*Producer),
	}
}
func (c *Producers) GetAll() map[string]*Producer {
	return c.m
}

func (c *Producers) Get(name string) (*Producer, bool) {
	v, ok := c.m[name]
	return v, ok
}

func (c *Producers) SetAll(Producers []*Producer)  {
	for _, v := range Producers {
		c.m[v.Name] = v
	}
}

func (c *Producers) Set(Producer *Producer) {
	c.m[Producer.Name] = Producer
}