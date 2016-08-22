package main

import (
     "fmt"
     . "github.com/mndrix/golog"
)

func main(){
m := NewMachine().Consult(`
  factor(hc, auxiliar, 1.8).
  factor(hc, asistente, 2.3).
  factor(hc, asociado, 2.7).
  factor(hc, titular, 3).

  factor(mto, auxiliar, 1.06).
  factor(mto, asistente, 1.69).
  factor(mto, asociado, 1.9).
  factor(mto, titular, 2.11).

  factor(tco, auxiliar, 2.12).
  factor(tco, asistente, 3.38).
  factor(tco, asociado, 3.8).
  factor(tco, titular, 4.22).

  valor(12120).
  factor_descuento_salud(0.04).
  factor_descuento_pension(0.04).

  valor_contrato(X,Y):-categoria(X,C),vinculacion(X,V),factor(V,C,F),valor(W),horas(X,H),Y is W * F * H.

  valor_pago(X,P):-valor_contrato(X,Y), duracion_contrato(X,D), P is Y / D.

  valor_descuento_salud(X,S):-valor_pago(X,P),factor_descuento_salud(F),S is P * F.
  valor_descuento_pension(X,S):-valor_pago(X,P),factor_descuento_pension(F),S is P * F.

  valor_pago_bruto(X,B):-valor_pago(X,V),valor_descuento_salud(X,S),valor_descuento_pension(X,P), B is V - (S + P).

  valores_finales(X,L):-valor_pago(X,V),valor_descuento_salud(X,S),valor_descuento_pension(X,P),valor_pago_bruto(X,B),L = [V,S,P,B].

  categoria(alejo, asociado).
  categoria(marce, asistente).
  vinculacion(alejo, tco).
  vinculacion(marce, hc).
  horas(alejo, 720).
  horas(marce, 180).
  duracion_contrato(alejo,5).
  duracion_contrato(marce,5).


`)

contratos := m.ProveAll(`valor_contrato(X,Y).`)
for _, solution := range contratos {
    fmt.Printf("%s contrato -> %s \n", solution.ByName_("X"), solution.ByName_("Y"))
}

pagos_brutos := m.ProveAll(`valor_pago(X,Y).`)
for _, solution := range pagos_brutos {
    fmt.Printf("%s pago -> %s \n", solution.ByName_("X"), solution.ByName_("Y"))
}

pagos_salud := m.ProveAll(`valor_descuento_salud(X,Y).`)
for _, solution := range pagos_salud {
    fmt.Printf("%s pago salud -> %s \n", solution.ByName_("X"), solution.ByName_("Y"))
}

pagos_pension := m.ProveAll(`valor_descuento_pension(X,Y).`)
for _, solution := range pagos_pension {
    fmt.Printf("%s pago pension -> %s \n", solution.ByName_("X"), solution.ByName_("Y"))
}

pagos_bruto := m.ProveAll(`valor_pago_bruto(X,Y).`)
for _, solution := range pagos_bruto {
    fmt.Printf("%s pago bruto -> %s \n", solution.ByName_("X"), solution.ByName_("Y"))
}

pagos_finales := m.ProveAll(`valores_finales(X,Y).`)
for _, solution := range pagos_finales {
    fmt.Printf("%s lista -> %s \n", solution.ByName_("X"), solution.ByName_("Y"))
}
}
