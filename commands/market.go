package Command

import (
  "github.com/jessehorne/tenland/data"
  "github.com/jessehorne/tenland/game"
)

func MarketCommandHandler(cmd []string, session *Game.Session) {
  session.Conn.Write([]byte("This command is a work-in-progress and isn't usable yet! Check back soon."));

  session.Conn.Write([]byte("\n" + Data.Cursor))
}

func NewMarketCommand() CommandType {
  hc := NewCommand("market", "'market <arg> <etc>' - Interact with the global player-ran market.")
  hc.Handler = MarketCommandHandler
  AllCommandsBig["market"] = `
Usage: 'market <arg> <etc>'
Interact with the global player-ran market.

In the market, you can buy and sell items from players no matter where you are
in Tenland (through magical couriers). You can sell any item that is in your
inventory.

Commands
========
market list\t\tList all items currently in the market.
market buy <num>\t\tBuy an item.
market sell <num> <amt>\t\tSell an item for a certain amount.
  `

  return hc
}
