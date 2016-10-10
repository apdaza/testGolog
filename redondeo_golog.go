package main

import (
     "fmt"
     . "github.com/mndrix/golog"
     //"strings"
)

func main(){
m := NewMachine().Consult(`
   residuo(N,R):-N @< 1, R is N.
   residuo(N,R):-N @>= 1, N1 is N - 1, residuo(N1,R).
   redondeo(N,R):-residuo(N,X),X @>= 0.5,R is N - X + 1.
   redondeo(N,R):-residuo(N,X),X @< 0.5,R is N - X.
`)

solutions := m.ProveAll(`redondeo(320.4, X).`)
for _, solution := range solutions {
    fmt.Printf("%s \n", solution.ByName_("X"))
}
}
