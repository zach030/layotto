package channel

import (
	"errors"
	"mosn.io/pkg/buffer"
	"sync"
)

func NewPipe() (*Pipe, *Pipe) {
	l1, l2 := &sync.Mutex{}, &sync.Mutex{}
	c1, c2 := sync.NewCond(l1), sync.NewCond(l2)
	b1, b2 := buffer.GetIoBuffer(4*1024), buffer.GetIoBuffer(4*1024)

	return &Pipe{
			l1: l1,
			b1: b1,
			c1: c1,

			l2: l2,
			b2: b2,
			c2: c2,

			lockOrder: true,
		},
		&Pipe{
			l1: l2,
			b1: b2,
			c1: c2,

			l2: l1,
			b2: b1,
			c2: c1,

			lockOrder: false,
		}
}

type Pipe struct {
	l1 *sync.Mutex
	c1 *sync.Cond
	b1 buffer.IoBuffer

	l2 *sync.Mutex
	c2 *sync.Cond
	b2 buffer.IoBuffer

	err       error
	lockOrder bool
}

func (p *Pipe) Read(b []byte) (int, error) {
	p.l1.Lock()
	for p.b1.Len() == 0 && p.err == nil {
		p.c1.Wait()
	}
	if err := p.err; err != nil {
		p.l1.Unlock()
		return 0, err
	}

	n, err := p.b1.Read(b)
	p.l1.Unlock()
	return n, err
}

func (p *Pipe) Write(b []byte) (int, error) {
	p.l2.Lock()
	if err := p.err; err != nil {
		p.l2.Unlock()
		return 0, err
	}
	p.b2.Write(b)
	p.l2.Unlock()
	p.c2.Signal()
	return len(b), nil
}

func (p *Pipe) Close() error {
	if p.lockOrder {
		p.l1.Lock()
		p.l2.Lock()
		p.err = errors.New("error closed")
		p.c1.Signal()
		p.c2.Signal()
		p.l2.Unlock()
		p.l1.Unlock()
	} else {
		p.l2.Lock()
		p.l1.Lock()
		p.err = errors.New("error closed")
		p.c2.Signal()
		p.c1.Signal()
		p.l1.Unlock()
		p.l2.Unlock()
	}
	return p.err
}
