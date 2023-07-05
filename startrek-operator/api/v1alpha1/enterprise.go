package v1alpha1

import (
	types "github.com/zerbitx/scifi-operators/commons/api/v1alpha1"
	"github.com/zerbitx/scifi-operators/commons/pkg"
)

func Enterprise() *types.Ship {
	return pkg.MakeShip("Enterprise", "Galaxy")
}

func UseUseless() {
	pkg.Useless()
}
