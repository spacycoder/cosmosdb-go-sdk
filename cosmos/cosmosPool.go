package cosmos

import (
	"bytes"
	"sync"
)

type cosmosPool struct {
	p *sync.Pool
}

func newCosmosPool() *cosmosPool {
	return &cosmosPool{
		&sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}

func (cp *cosmosPool) Get() *bytes.Buffer {
	return cp.p.Get().(*bytes.Buffer)
}

func (cp *cosmosPool) Put(buffer *bytes.Buffer) {
	buffer.Reset()
	cp.p.Put(buffer)
}

var buffers = newCosmosPool()
