package player

import (
  "time"
  "math/rand"
  "github.com/inazak/reversi/point"
  "github.com/inazak/reversi/game"
)

func init() {
  rand.Seed(time.Now().UnixNano())
}

type RandomPlayer struct {
}

func NewRandomPlayer() RandomPlayer {
  return RandomPlayer{}
}

func (r RandomPlayer) Play(g game.Game) (p point.Point, pass, giveup bool) {
  ps := g.GetPuttablePoint(g.GetTurn())
  p   = ps[ rand.Intn(len(ps)) ]
  return p, false, false
}


