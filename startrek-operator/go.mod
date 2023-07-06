module github.com/zerbitx/scifi-operators/startrek-operator

go 1.18

require (
	github.com/zerbitx/scifi-operators/commons v0.0.3
	github.com/zerbitx/scifi-operators/commons/api v0.0.1
	github.com/zerbitx/scifi-operators/startrek-operator/api v0.0.4
)

replace github.com/zerbitx/scifi-operators/startrek-operator/api => ./api
