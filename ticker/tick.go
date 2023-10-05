package ticker

import (
	"fmt"
	"net/http"
	"time"
)

const UrlsSelf = "http://127.0.0.1:3000/api/tick/update"

type Tick struct {
	t *time.Ticker
}

func NewTick() *Tick {
	return &Tick{
		t: time.NewTicker(1 * time.Minute),
	}
}

func newReq() error {
	_, err := http.Get(UrlsSelf)
	if err != nil {
		return err
	}
	return nil
}

func (t *Tick) Loop() {
	go func() {
		for {
			select {
			case <-t.t.C:
				fmt.Println("Tick")
				if err := newReq(); err != nil {
					fmt.Println(err)
				}
			}
		}
	}()
}
