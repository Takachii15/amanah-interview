package optional

import "fmt"

func Test(){
  var first string
  var seconf string
  fmt.Println("Nomor rahasia: ")
  fmt.Scanln(&first)
  fmt.Println("Nomor tebakan: ")
  fmt.Scanln(&seconf)

  truth := 0
  false := 0

  for i := 0; i<len(first); i++{
    if first[i] != seconf[i] {
      truth = truth + 1
    } else {
      false = false + 1 }
  }

  fmt.Printf("alhamdulillah %d; subhanallah %d" , truth, false )
}
