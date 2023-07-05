package main

import (
	"fmt"

	st "github.com/zerbitx/scifi-operators/startrek-operator/pkg"
	sw "github.com/zerbitx/scifi-operators/starwars-operator/pkg"
)

func main() {
	ent := st.GetShip()
	xwing := sw.GetShip()

	fmt.Println("Enterprise: ", ent.Name, ent.Class)
	fmt.Println("X-Wing: ", xwing.Name, xwing.Class)
}
