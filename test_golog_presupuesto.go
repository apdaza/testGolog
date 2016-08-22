package main

import (
     "fmt"
     . "github.com/mndrix/golog"
)

func main(){
m := NewMachine().Consult(`
  rubro(11,300000).
  rubro(12,20000).
  rubro(13,100000).
  rubro(14,1300000).
  rubro(15,100).
  rubro(16,140000).
  rubro(17,250000).
  rubro(18,3000000).

  cdp(1,11,30000).
  cdp(2,12,2000).
  cdp(3,13,10000).
  cdp(4,14,130000).
  cdp(5,15,100000).
  cdp(6,16,14000).
  cdp(7,17,25000).
  cdp(8,18,300000).

  rp(121,1,20).
  rp(122,2,20).
  rp(123,3,20000).
  rp(124,5,200).

  validar_cdp(X,R,Y):-cdp(X,R,V),rubro(R,D),V @=< D,Y is 1.
  validar_cdp(X,R,Y):-cdp(X,R,V),rubro(R,D),V @> D,Y is 2.

  validar_rp(X,C,Y):-rp(X,C,V),cdp(C,R,D),V @=< D,validar_cdp(C,R,Z),Z == 1,Y is 1.
  validar_rp(X,C,Y):-rp(X,C,V),cdp(C,R,D),V @=< D,validar_cdp(C,R,Z),Z == 2,Y is 2.
  validar_rp(X,C,Y):-rp(X,C,V),cdp(C,R,D),V @> D,validar_cdp(C,R,Z),Z == 1,Y is 2.
  validar_rp(X,C,Y):-rp(X,C,V),cdp(C,R,D),V @> D,validar_cdp(C,R,Z),Z == 2,Y is 2.
`)

cdps := m.ProveAll(`validar_cdp(X,R,Y).`)
for _, solution := range cdps {
  fmt.Printf("El CDP No. %s tiene Rubro -> %s y es %s \n", solution.ByName_("X"), solution.ByName_("R"), solution.ByName_("Y"))
}

rps := m.ProveAll(`validar_rp(X,R,Y).`)
for _, solution := range rps {
  fmt.Printf("El RP No. %s tiene CDP -> %s y es %s \n", solution.ByName_("X"), solution.ByName_("R"), solution.ByName_("Y"))
}
/*

*/
}
