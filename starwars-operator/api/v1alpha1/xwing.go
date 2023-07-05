package v1alpha1

import (
	"github.com/zerbitx/scifi-operators/commons/api/v1alpha1"
	"github.com/zerbitx/scifi-operators/commons/pkg"
	"github.com/zerbitx/scifi-operators/commons/pkg/utils"
)

func XWing() *v1alpha1.Ship {
	return pkg.MakeShip("X-Wing", "Starfighter")
}

func UseUseless() {
	utils.Blah("test")
}
