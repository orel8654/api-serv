package ticker

import (
	"fmt"
	"time"
)

const UrlsSelf = "http://127.0.0.1:3000/api/tick/update"

type Tick struct {
	t *time.Ticker
}

func NewTick(tk time.Duration) *Tick {
	return &Tick{
		t: time.NewTicker(tk),
	}
}

func (t *Tick) LoopAccept(fn func() error) {
	go func() {
		for {
			select {
			case <-t.t.C:
				fmt.Println("Tick")
				if err := fn(); err != nil {
					fmt.Println(err)
				}
			}
		}
	}()
}
