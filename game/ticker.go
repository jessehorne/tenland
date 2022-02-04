package Game

import (
  "time"
)

var Ticker = time.NewTicker(time.Second)
var Done = make(chan bool)
var TickCount = 0

func StartTicker() {
  go TickerLoop()
}

func TickerLoop() {
  for {
    select {
    case <- Done:
      return
    case <-Ticker.C:
      TickCount++
    }
  }
}
