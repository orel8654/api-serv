package ticker

import "time"

func Ticker() {
	ticker := time.NewTicker(5 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				//HandlerUpdateCurrency(configDB, configAPI)
			}
		}
	}()
	defer func() {
		ticker.Stop()
		done <- true
	}()
}
