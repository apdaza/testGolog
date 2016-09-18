package main

import (
     "fmt"
     . "github.com/mndrix/golog"
     //"strings"
)

func main(){
m := NewMachine().Consult(`
  factor(hc, auxiliar, 1.8, 2016).
  factor(hc, asistente, 2.3, 2016).
  factor(hc, asociado, 2.7, 2016).
  factor(hc, titular, 3, 2016).

  factor(mto, auxiliar, 1.06, 2016).
  factor(mto, asistente, 1.69, 2016).
  factor(mto, asociado, 1.9, 2016).
  factor(mto, titular, 2.11, 2016).

  factor(tco, auxiliar, 2.12, 2016).
  factor(tco, asistente, 3.38, 2016).
  factor(tco, asociado, 3.8, 2016).
  factor(tco, titular, 4.22, 2016).

  valor(12120, 2016).
  factor_descuento_salud(0.04, 2016).
  factor_descuento_pension(0.04, 2016).

  valor_contrato(X,P,Y):-categoria(X,C,P),vinculacion(X,V,P),factor(V,C,F,P),valor(W,P),horas(X,H,P),Y is W * F * H.

  valor_pago(X,V,P):-valor_contrato(X,V,Y), duracion_contrato(X,D,V), P is Y / D.

  valor_descuento_salud(X,V,S):-valor_pago(X,V,P),factor_descuento_salud(F,V),S is P * F.
  valor_descuento_pension(X,V,S):-valor_pago(X,V,P),factor_descuento_pension(F,V),S is P * F.

  valor_pago_bruto(X,P,B):-valor_pago(X,P,V),valor_descuento_salud(X,P,S),valor_descuento_pension(X,P,D), B is V - (S + D).

  valores_finales(X,L,P):-valor_pago(X,P,V),valor_descuento_salud(X,P,S),valor_descuento_pension(X,P,D),valor_pago_bruto(X,P,B),L = [V,S,D,B].

  categoria(alejo, asociado, 2016).
  categoria(marce, asistente, 2016).
  vinculacion(alejo, tco, 2016).
  vinculacion(marce, hc, 2016).
  horas(alejo, 720, 2016).
  horas(marce, 180, 2016).
  duracion_contrato(alejo,5, 2016).
  duracion_contrato(marce,5, 2016).

  factores(X,T,P,L):-findall((X, Y, N, Z),(factor(X,T,Y,N,Z,P)),L).

  valores(X,T,P,L):-findall((X, Y, N, Z, R),((factor(X,T,Y,N,Z,P),Y==porcentaje,valor_pago(X,P,V),R is P * Z)),L).
  valores(X,T,P,L):-findall((X, Y, N, Z, R),((factor(X,T,Y,N,Z,P),Y==fijo,R is Z)),L).

  factor(alejo, descuento, porcentaje, salud, 0.04, 2016).
  factor(alejo, descuento, porcentaje, pension, 0.04, 2016).
  factor(alejo, descuento, porcentaje, solidaridad, 0.01, 2016).
  factor(alejo, descuento, fijo, fondo, 20000, 2016).
  factor(alejo, descuento, fijo, prestamo, 200000, 2016).
  factor(alejo, descuento, fijo, sindicato, 10000, 2016).

  factor(marce, descuento, porcentaje, salud, 0.04, 2016).
  factor(marce, descuento, porcentaje, pension, 0.04, 2016).
  factor(marce, descuento, porcentaje, solidaridad, 0.01, 2016).
  factor(marce, descuento, fijo, fondo, 1000, 2016).
  factor(marce, descuento, fijo, prestamo, 5000, 2016).
  factor(marce, descuento, fijo, sindicato, 3000, 2016).
`)
/**
* factor(docente, descuento|suma, porcentaje|fijo, nombre del factor, valor del factor, periodo)
*/
descuentos := m.ProveAll(`valores(alejo,descuento,2016,L).`)
for _, solution := range descuentos {
    //fmt.Printf("%s -> %s -> %s -> %s\n", solution.ByName_("X"), solution.ByName_("Y"), solution.ByName_("N"), solution.ByName_("Z"))
    fmt.Printf("%s",solution.ByName_("L"))
}

/*contratos := m.ProveAll(`valor_contrato(X,Y).`)
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
}*/

/*pagos_finales := m.ProveAll(`valores_finales(X,Y,2016).`)
for _, solution := range pagos_finales {
    fmt.Printf("%s lista -> %s \n", solution.ByName_("X"), solution.ByName_("Y"))
    r := strings.NewReplacer("(", "",")", "","[]", "","'.'", "",)
    cad := r.Replace(solution.ByName_("Y").String())
    fmt.Println(cad)
    pair := strings.Split(cad, ",")
    fmt.Println(pair[1])
}*/

}
