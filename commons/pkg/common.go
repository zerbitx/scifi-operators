package pkg

import (
	"github.com/zerbitx/scifi-operators/commons/api/v1alpha1"
)

func MakeShip(name, class string) *v1alpha1.Ship {
	return &v1alpha1.Ship{
		Name:  name,
		Class: class,
	}
}

func Useless() {
	// Do nothing
}
