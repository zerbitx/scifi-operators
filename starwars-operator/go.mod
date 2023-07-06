module github.com/zerbitx/scifi-operators/starwars-operator

go 1.18

require (
	github.com/zerbitx/scifi-operators/commons/api v0.0.3
	github.com/zerbitx/scifi-operators/starwars-operator/api v0.0.3
)

require github.com/zerbitx/scifi-operators/commons v0.0.3 // indirect

replace github.com/zerbitx/scifi-operators/starwars-operator/api => ./api
