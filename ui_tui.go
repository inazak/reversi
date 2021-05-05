package ui

import (
  "github.com/inazak/reversi/point"
  "github.com/inazak/reversi/stone"
  "github.com/inazak/reversi/game"
  "github.com/nsf/termbox-go"
)

const (
  fgColor = termbox.ColorWhite
  bgColor = termbox.ColorBlack
  fgEmColor = termbox.ColorBlack
  bgEmColor = termbox.ColorWhite
)

var inputmode  bool
var inputpoint []point.Point
var inputpointindex int

var display = []string{}
var display6x6 = []string{
        //01234567890123456789012345678901234567890123456789
/*  0 */ "                   Reversi                        ",
/*  1 */ "                                                  ",
/*  2 */ "    1 2 3 4 5 6            Turn:   XXXXX          ",
/*  3 */ "  a _ _ _ _ _ _                                   ",
/*  4 */ "  b _ _ _ _ _ _            Black:  XX             ",
/*  5 */ "  c _ _ _ _ _ _            White:  XX             ",
/*  6 */ "  d _ _ _ _ _ _                                   ",
/*  7 */ "  e _ _ _ _ _ _            History:  XXXXX[1a]    ",
/*  8 */ "  f _ _ _ _ _ _                      XXXXX[1a]    ",
/*  9 */ "                                     XXXXX[1a]    ",
/* 10 */ "                                     XXXXX[1a]    ",
/* 11 */ "                                     XXXXX[1a]    ",
/* 12 */ "                                                  ",
/* 13 */ "  MESSAGE AREA XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX  ",
/* 14 */ "                                                  ",
}

type Coord struct {
  X int
  Y int
}

type TUIController struct {}

func NewTUIController() TUIController {
  return TUIController{}
}

func (c TUIController) Init(g game.Game) {

  // termbox init and event-loop
  err := termbox.Init()
  if err != nil {
    panic(err)
  }
  defer termbox.Close()

  switch g.GetBoardSize() {
  case 6:
    display = display6x6
  default:
    panic
  }

  eventQueue := make(chan termbox.Event)
  go func(){
    for {
      eventQueue <- termbox.PollEvent()
    }
  }()

  render()

  go func() {
    for {
      select {
      case ev := <-eventQueue:
        if ev.Type == termbox.EventKey {
          switch {
          case ev.Ch == 'p':
            if inputmode {
              render()
            }
          case ev.Key == termbox.KeyArrowUp:
            if inputmode {
              inputpointindex += 1
              if inputpointindex >= len(inputpoint) {
                inputpointindex = 0
              }
              c := pointToCoord(inputpoint[inputpointindex])
              setText(c.X, c.Y, bgColor, fgColor, "_")
              render()
            }
          case ev.Key == termbox.KeyEsc:
            if inputmode {
              render()
            }
          }
        }
      }
    }
  }()
}

func (c TUIController) Gameset(g game.Game, winner stone.Stone) {
}

func (c TUIController) Wait(g game.Game)  {
}

func (c TUIController) Pass(g game.Game) {
}

func (c TUIController) Giveup(g game.Game) {
}

func (c TUIController) Put(g game.Game, p point.Point) {
}

func (c TUIController) Input(g game.Game) (p point.Point, pass, giveup bool) {

  //input point or pass, giveup
  inputpoint := g.GetPuttablePoint(g.GetTurn())
  inputpointindex = 0

  coord := pointToCoord(inputpoint[0])
  setText(coord.X, coord.Y, bgColor, fgColor, "_")
  render()

  inputmode = true

  return point.Point{}, false, false
}

func pointToCoord(p point.Point) Coord {

  // "1a" => Point{ X: 0, Y: 0 } => Coord{ X: 4, Y: 3 }
  // "2b" => Point{ X: 1, Y: 1 } => Coord{ X: 6, Y: 4 }

  c := Coord{ X: p.X * 2 + 4, Y: p.Y + 3 }
  return c
}

/*
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


func pointToFormat1A(p point.Point) (s, xs, ys string) {
  xs = string(byte(p.X + 0x31))
  ys = string(byte(p.Y + 0x61))
  s  = xs + ys
  return s, xs, ys
}

// Point{ X: 0, Y: 0 } => "1a"
// Point{ X: 1, Y: 1 } => "2b"
func format1AtoPoint(s string) (point.Point, error) {
  p := point.Unavailable
  if len(s) != 2 {
    return p, fmt.Errorf("unexpected format")
  }
  if s[0] < '0' || '8' < s[0] { //LIMIT 8x8
    return p, fmt.Errorf("unexpected format")
  }
  if s[1] < 'a' || 'h' < s[1] { //LIMIT 8x8
    return p, fmt.Errorf("unexpected format")
  }

  p.X = int(s[0] - 0x31) // '1' => 0
  p.Y = int(s[1] - 0x61) // 'a' => 0

  return p, nil
}
*/

func render() {
  termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

  //title
  setText(0, 0, fgEmColor, bgEmColor, display[0])

  //other
  for i:=1; i<len(display); i++ {
    setText(0, i, fgColor, bgColor, display[i])
  }

  //reflesh
  termbox.Flush()
}

func setText(x, y int, fg, bg termbox.Attribute, msg string) {
  for _, c := range msg {
	  termbox.SetCell(x, y, c, fg, bg)
    x++
  }
}

