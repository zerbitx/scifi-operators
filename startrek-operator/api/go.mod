module github.com/zerbitx/scifi-operators/startrek-operator/api

go 1.18

require (
	github.com/zerbitx/scifi-operators/commons v0.0.3
	github.com/zerbitx/scifi-operators/commons/api v0.0.3
)

replace (
	github.com/zerbitx/scifi-operators/commons => ../../commons
	github.com/zerbitx/scifi-operators/commons/api => ../../commons/api
)
