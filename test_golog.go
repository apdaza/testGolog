package main

import (
     "fmt"
     . "github.com/mndrix/golog"
)

func main(){
m := NewMachine().Consult(`
    father(john).
    father(jacob).

    mother(sue).

    fatherof(john, maria).
    fatherof(jean, juan).
    fatherof(juan, ana).
    fatherof(john, mercedes).
    fatherof(ana, pedro).

    parent(X) :-
        father(X).
    parent(X) :-
        mother(X).
    brotherof(X, Y) :-
        fatherof(C, X), fatherof(C, Y).
`)

if m.CanProve(`father(john).`) {
    fmt.Printf("john is a father\n")
}

solutions := m.ProveAll(`parent(X).`)
for _, solution := range solutions {
    fmt.Printf("%s is a parent\n", solution.ByName_("X"))
}

solutions2 := m.ProveAll(`brotherof(X, maria).`)
for _, solution2 := range solutions2 {
    fmt.Printf("%s is a brother of mercedes \n", solution2.ByName_("X"))
}
}
