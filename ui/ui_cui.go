package ui

import (
  "fmt"
  "github.com/inazak/reversi/point"
  "github.com/inazak/reversi/stone"
  "github.com/inazak/reversi/game"
)

type CUIController struct {}

func NewCUIController() CUIController {
  return CUIController{}
}

func (c CUIController) Init(g game.Game) {
  printInfo(g)
}
func (c CUIController) Gameset(g game.Game) {
}
func (c CUIController) Wait(g game.Game)  {
  fmt.Printf("waiting %v move ...\n", g.GetTurn())
  fmt.Printf("\n")
}
func (c CUIController) Pass(g game.Game) {
}
func (c CUIController) Giveup(g game.Game) {
}
func (c CUIController) Put(g game.Game, p point.Point) {
  printInfo(g)
}
func (c CUIController) Input(g game.Game) (p point.Point, pass, giveup bool) {
  var x, y int
  fmt.Printf("input move x,y: ")
  fmt.Scanf("%d,%d\n", &x, &y)

  return point.Point{X:x, Y:y}, false, false
}


func printInfo(g game.Game) {
  fmt.Printf("\n")
  fmt.Printf("Turn: %v\n", g.GetTurn())
  fmt.Printf("\n")

  for y:=0; y<g.GetBoardSize(); y++ {
    fmt.Printf("  ")
    for x:=0; x<g.GetBoardSize(); x++ {
      m := stoneToMark(g.GetStone(point.Point{X:x, Y:y}))
      fmt.Printf("%v ", m)
    }
    fmt.Printf("\n")
  }
  fmt.Printf("\n")
}

func stoneToMark(s stone.Stone) string {
  switch s {
  case stone.None:
    return "_"
  case stone.Black:
    return "x"
  case stone.White:
    return "o"
  default:
    panic("stone type is out of range")
  }
}

