package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type PieceType int
type Color int

const (
	_ Color = iota
	White
	Black
)

const (
	_ PieceType = iota
	King
	Queen
	Pawn
	Bishop
	Knight
	Rook
	Blank
)

var Row = map[int]string{
	7: `h`,
	6: `g`,
	5: `f`,
	4: `e`,
	3: `d`,
	2: `c`,
	1: `b`,
	0: `a`,
}

var RowRev = map[string]int{
	`h`: 7,
	`g`: 6,
	`f`: 5,
	`e`: 4,
	`d`: 3,
	`c`: 2,
	`b`: 1,
	`a`: 0,
}

type Piece struct {
	Type  PieceType
	Color Color
}

func (p *Position) makeMove(oldS, newS string) {
	newRow, _ := strconv.Atoi(string(newS[1]))
	oldRow, _ := strconv.Atoi(string(oldS[1]))

	p[newRow-1][RowRev[string(newS[0])]] = p[oldRow-1][RowRev[string(oldS[0])]]

	p[oldRow-1][RowRev[string(oldS[0])]] = &Piece{Type: Blank}
}

func (p *Piece) String() string {
	var s string
	handleCase := func(s string, c Color) string {
		switch c {
		case White:
			return strings.ToUpper(s)
		default:
			return strings.ToLower(s)
		}
	}
	switch p.Type {
	case Queen:
		s = handleCase(`q`, p.Color)
	case King:
		s = handleCase(`k`, p.Color)
	case Pawn:
		s = handleCase(`p`, p.Color)
	case Knight:
		s = handleCase(`n`, p.Color)
	case Rook:
		s = handleCase(`r`, p.Color)
	case Bishop:
		s = handleCase(`b`, p.Color)
	default:
		s = ` `
	}
	return s
}

type Position [8][8]*Piece

func (p *Position) setPiece(row, col int, kind PieceType, c Color) {
	p[row][col] = &Piece{Type: kind, Color: c}
}

type Game struct {
	CurrentPosition *Position
}

var G *Game

func InitGame() {
	G = new(Game)
	G.CurrentPosition = new(Position)
	G.CurrentPosition.GetStartPosition()
}

func (p *Position) String() string {
	var res string = "  ---------------------------------\n"
	for i := len(p) - 1; i >= 0; i-- {
		res = res + `|` + fmt.Sprint(i+1) + `|`
		for _, v := range p[i] {
			res += fmt.Sprintf(` %s |`, v)
		}
		res += "\n  ---------------------------------\n"
	}
	res += "| | a | b | c | d | e | f | g | h |\n  ---------------------------------"
	return res
}

func setPieceOnBoard(pos *Position, row, column int, t PieceType, c Color) {
	pos[row][column] = &Piece{Type: t, Color: c}
}

func (position *Position) GetStartPosition() {
	position.setPiece(0, 0, Rook, White)
	position.setPiece(0, 7, Rook, White)
	position.setPiece(7, 0, Rook, Black)
	position.setPiece(7, 7, Rook, Black)
	position.setPiece(0, 1, Knight, White)
	position.setPiece(0, 6, Knight, White)
	position.setPiece(7, 1, Knight, Black)
	position.setPiece(7, 6, Knight, Black)
	position.setPiece(0, 2, Bishop, White)
	position.setPiece(0, 5, Bishop, White)
	position.setPiece(7, 2, Bishop, Black)
	position.setPiece(7, 5, Bishop, Black)
	position.setPiece(0, 3, Queen, White)
	position.setPiece(0, 4, King, White)
	position.setPiece(7, 3, Queen, Black)
	position.setPiece(7, 4, King, Black)
	for i := 0; i < 8; i++ {
		position.setPiece(1, i, Pawn, White)
		position.setPiece(6, i, Pawn, Black)
	}
	for i1, v := range position {
		for i2, v1 := range v {
			if v1 == nil {
				position[i1][i2] = &Piece{Type: Blank}
			}
		}
	}
}

func NewGame() *Game {
	return nil
}

func main() {
	InitGame()
	fmt.Println(G.CurrentPosition)

	G.CurrentPosition.makeMove(`e2`, `e4`)
	time.Sleep(time.Second * 3)
	fmt.Println(G.CurrentPosition)
}
