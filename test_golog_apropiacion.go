package main

import (
     "fmt"
     . "github.com/mndrix/golog"
)

func main(){
m := NewMachine().Consult(`
  suma_elementos([], 0).
  suma_elementos([X|Xs], S):-suma_elementos(Xs, S2),S is S2 + X.
  verifica_hijos(P,[],R):-R is 1.
  verifica_hijos(P,H,R):-suma_elementos(P, S1),suma_elementos(H, S2),S2 @=< S1,R is 1.
  verifica_hijos(P,H,R):-suma_elementos(P, S1),suma_elementos(H, S2),S2 @> S1,R is 2.
  ingresos_egresos(I,E,R):-suma_elementos(I, S1),suma_elementos(E, S2),S2 == S1,R is 1.
  ingresos_egresos(I,E,R):-suma_elementos(I, S1),suma_elementos(E, S2),S2 \== S1,R is 2.
  `)

resultados := m.ProveAll(`suma_elementos([200,500,700,600,20],Y).`)
for _, solution := range resultados {
    fmt.Printf("suma is: %s \n", solution.ByName_("Y"))
}
resultados2 := m.ProveAll(`verifica_hijos([200,500,700,600,20],[2000000,3],Y).`)
for _, solution := range resultados2 {
    fmt.Printf("hijos is: %s \n", solution.ByName_("Y"))
}

resultados3 := m.ProveAll(`ingresos_egresos([200,500,700,600,20],[2000000,3],Y).`)
for _, solution := range resultados3 {
    fmt.Printf("ingresos vs egresos is: %s \n", solution.ByName_("Y"))
}

}
