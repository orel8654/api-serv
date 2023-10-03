package ticker

import (
	"fmt"
	"time"
)

type Tick struct {
	t *time.Ticker
}

func NewTick() *Tick {
	return &Tick{
		t: time.NewTicker(1 * time.Minute),
	}
}

func (t *Tick) Loop() {
	go func() {
		for {
			select {
			case <-t.t.C:
				fmt.Println("Tick")
			}
		}
	}()
}
