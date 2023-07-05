package pkg

import (
	commons "github.com/zerbitx/scifi-operators/commons/api/v1alpha1"
	"github.com/zerbitx/scifi-operators/starwars-operator/api/v1alpha1"
)

func GetShip() *commons.Ship {
	return v1alpha1.XWing()
}

func Useless() {}
