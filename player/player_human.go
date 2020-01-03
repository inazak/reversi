package player

import (
  "github.com/inazak/reversi/point"
  "github.com/inazak/reversi/game"
  "github.com/inazak/reversi/ui"
)

type HumanPlayer struct {
  ctrl  ui.Controller
}

func NewHumanPlayer(c ui.Controller) HumanPlayer {
  return HumanPlayer{ ctrl: c }
}

func (r HumanPlayer) Play(g game.Game) (p point.Point, pass, giveup bool) {
  p, pass, giveup = r.ctrl.Input(g)
  return p, pass, giveup
}


