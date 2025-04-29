package main

import(
  "fmt"
  "os"
  "math/rand"
  "strconv"
)

func main() {
  args := os.Args
  numArgs := len(os.Args)
  
  if numArgs != 4 {
    fmt.Println("Takes Three Arguments")
    fmt.Println("Argument One: Number of Dice")
    fmt.Println("Argument Two: Number of Sides on the Dice")
    fmt.Println("Argument Three: Bonus to Total Roll")
    os.Exit(0)
  }

  numDice := args[1]
  sideDice := args[2]
  numBonus := args[3]

  diceRoller(numDice, sideDice, numBonus)
}

func diceRoller(Num string, Side string, Bonus string){
  fmt.Printf("There is %v number of dice with %v number of side\n", Num, Side)
  NumInt, _ := strconv.Atoi(Num)
  SideInt, _ := strconv.Atoi(Side)
  BonusInt, _ := strconv.Atoi(Bonus)
  for Count:=1; Count <= NumInt; Count++ {
    Results := rand.Intn(SideInt) + 1 + BonusInt
    fmt.Printf("Die %v: %v\n", Count, Results)
  }
}
