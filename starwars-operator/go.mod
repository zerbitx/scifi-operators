module github.com/zerbitx/scifi-operators/starwars-operator

go 1.18

replace github.com/zerbitx/scifi-operators/starwars-operator/api => ./api

require (
	github.com/zerbitx/scifi-operators/commons/api v0.0.1
	github.com/zerbitx/scifi-operators/starwars-operator/api v0.0.0-00010101000000-000000000000
)

require github.com/zerbitx/scifi-operators/commons v0.0.0-20230705203653-03474eb64713 // indirect
