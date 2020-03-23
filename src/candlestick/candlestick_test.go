package main
import (
  "testing"
  "fmt"
)

type testtable struct{
  question []int
  expected int
}

func TestUniques(t *testing.T) {
  tab := []testtable{
    testtable{ []int{3, 6, 5, 6, 4, 5}, 4} ,
    testtable{ []int{3, 6, 5, 7, 4, 5}, 3} ,
    testtable{ []int{13, 6, 5, 1, 4, 5}, 1} ,
    testtable{ []int{3, 13, 13, 1, 4, 5}, 6} ,
  }
  
  for i := 0; i < len(tab); i++ {
    r := DecideTrick(tab[i].question)
    if r != tab[i].expected {
      t.Fail()
      fmt.Printf("Test : %s, expected %d, got %d \n", tab[i].question, tab[i].expected, r)
    }
  }
}

type containsTest struct {
  qArr []int
  qCom int
  exp int
}

func TestContains(t *testing.T) {
  ar1 := []int{4, 7, 3, 4, 6, 6, 4}
  tab := []containsTest{
    containsTest{ar1, 4, 3},
    containsTest{ar1, 7, 1},
    containsTest{ar1, 3, 1},
    containsTest{ar1, 6, 2}, 
    containsTest{ar1, 15, 0},
  }
  
  for _, v := range tab {
    res := contains(v.qArr, v.qCom)
    if res != v.exp {
      t.Fail()
      fmt.Printf("%s, %d: expected %d, got %d\n", v.qArr, v.qCom, v.exp, res)
    }
  }
}
