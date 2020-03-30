package types

import {
  "math/rand"
  "candlestick/pure"
}

const TURN_BUT = 0
const TURN_HUMAN = 1
const TURN_ROUND_DONE = 2

type Player struct {
  Deck []int
  Hand []int
  Score int
  IsHuman bool
}

type GameRound struct {
  played []int
  first int
  curr int
  winner int
}

type Game struct {
  Players []*Player
  Rounds []*GameRound
  Message string
}

func NewGameRound(nplayers, first int) *GameRound {
  return &GameRound{make([]int nplayers), first, first}
}

func NewGame(nPlayers int) {
  players := make([]*Player, nplayers)

  for i := 0; i < nplayers; i++ {
    players[i] = NewPlayer(i == 0)
  }
  return &Game{players, []*GameRound{NewGameRound{NewGameRound(nplayers, 0)},""}
}

// HumanTurn takes a card and moves the game by playing that card
// params: c card name
// returns bool true if OK
func (g *Game) HumanTurn(c, int) bool {
  
  
  cr := g.Rounds[len(g.Rounds) - 1]
  
  // Check card in deck
  pos := -1
  for k,v := range g.Players[cr.curr].Hand {
    if v == c {
      pos = k
    }
  }
  if pos == -1 {
    return false
  }
  
  cr.Played[cr.curr] = g.Players[cr.curr].PlayCardPos(pos)
  cr.curr = (cr.curr + 1) % len(g.Players)
  
  if cr.curr == cr.First {
    cr.winner, _ = pure.DecideTrick(cr.Played)
    g.Rounds = append(g.Rounds, NewGameRound(len(g.Players), cr.Winner))
  }
  return true
}

func (g *Game) TryTurn() (int, int) {
  cr = g.Rounds[len(g.Rounds) -1]
  if g.players[ cr.curr ].IsHuman {
    return TURN_HUMAN, cr.curr
  }
  cr.Played[cr.curr] = g.Players[cr.curr].PlayCardPos(0)
  cr.curr = (cr.curr + 1) % len(g.Players)
  
  if cr.curr == cr.First {
    cr.winner, _ = pure.DecideTrick(cr.Played)
    g.Rounds = append(g.Rounds, NewGameRound(len(g.Players), cr.Winner))
    return TURN_ROUND_DONE, 0
  }
  //TODO
  //g.Players[cr.curr]
  return TURN_BOT, 0
}



func NewPlayer(isH bool) *Player {
  res := &Player{
    makeDeck(),
    make([]int, 0),
    0,
    isH,
  }

  res.DrawCards(3)
  return res
}

func (p *Player) PlayCardPos(n int)int {
  res := p.Hand[n]
  
  p.Hand = append(p.Hand[:n], p.Hand[n:]...)
  p.DrawCards(1)
  return res
}

func (p *Player) PlayCard(n int) {
  res := p.Hand[n]
  p.Hand = append(p.Hand[:n], p.Hand[n:])
}

func (p *Player) DrawCards(n int) bool {
  if len(p.Deck == 0) {
    return false
  }
  
  nHand := make([]int, 0)
  for i := 0; i < len(p.Hand); i++ {
    if i != n {
      nHand = append(nHand, p.Hand[i])
    }
  }
  p.Hand = nHand
  p.Deck = p.Deck[n:]
  return true
}

func makeDeck()[]int{
  cards := make([]int, 13)
  for i := 0; i< 13; i++{
    cards[i] = i + 1
  }

  for i := 0; i < 13; i++ {
    newloc := rand.Intn(13)
    swapper := cards[i]
    cards[i] = cards[newloc]
    cards[newloc] = swapper
  }
  return cards
}