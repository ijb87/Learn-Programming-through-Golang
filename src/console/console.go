package console

func (self *Player) DrawCards(n int) bool {
  if len(self.deck) == 0 {
    return false
  }
  self.hand = append(self.hand, self.deck[:n]...)
  self.deck = self.deck[n:]
  return true
}

func (self *Player) ChooseCard() int {
  a := self.hand[0]
  if self.isHuman {
    a = AskCard(self.hand)
  }

  h2 := make([]int, 0)
  for _, v := range self.hand {
    if v != a {
      h2 = append(h2, v)
    }
  }
  self.hand = h2
  self.DrawCards(1)
  return a
}

func AskCard(cards []int) int {
  mess := "Please choose a card from "
  for _, v := range cards {
    mess += strconv.Itoa(v) + ", "
  }
  for {
    a := asker.AskInt(mess)
    if contains(cards, a) > 0 {
      return a
    }
    fmt.Println("Stick to the list please")
  }
}

func round(pl []*Player, first int) int {
  pnum := first
  complete := false
  chosens := make([]int, len(pl))
  for !complete {
    chosens[pnum] = pl[pnum].ChooseCard()
    fmt.Printf("p%d : %d\n", pnum, chosens[pnum])
    pnum = (pnum + 1) % len(pl)
    if pnum == first {
      complete = true
    }
  }

  winP, winC := DecideTrick(chosens)
  pl[winP].score++
  fmt.Printf("Player %d wins with %d"\n", winP, winC)

  scoresMess := "Scores : "
  for _, v := range pl {
    scoresMess += strconv.Itoa(v.score) + " "
  }
  fmt.Println(scoresMess)

  return winP
}

func round(pl []*Player, first int) int {
  pnum := first
  complete := false
  chosens := make([]int, len(pl))
  for !complete {
    chosens[pnum] = pl[pnum].ChooseCard()
    fmt.Printf("p%d : %d\n", pnum, chosens[pnum])
    pnum = (pnum + 1) % len(pl)
    if pnum == first {
      complete = true
    }
  }

  winP, winC := DecideTrick(chosens)
  pl[winP].score++
  fmt.Printf("Player %d wins with %d"\n", winP, winC)

  scoresMess := "Scores : "
  for _, v := range pl {
    scoresMess += strconv.Itoa(v.score) + " "
  }
  fmt.Println(scoresMess)

  return winP
}

func playGame(nplayers int) int {
  if nplayers <= 0 {
    return -1
  }


  players := make([]*Player, nplayers)

  for i := 0; i < nplayers; i++ {
    players[i] = NewPlayer(i == 0)
    // fmt.Printf("%s, \n%s\n, %d\n\n", players[i].deck, players[i].hand, players[i].score)
  }

  w := -1
  for i := 0; i < 13; i ++ {
    if w == -1 {
      w = rand.Intn(nplayers)
    }
    w = round(players, w)
  }

  // Get winner.
  ws := -1
  wp := -1

  for k, p := range players {
    if p.score > ws {
      count := 0
      for i := 0; i < len(players); i++ {
        if p.score == players[i].score {
          count++
        }
      }
      if count == 1 {
        ws = p.score
        wp = k
      }
    }
}

if wp == -1 {
  fmt.Printf("Draw\n")
  return -1
}

fmt.Printf("Player %d wins\n\n", wp)
return wp
}

