package cgin

type IClass interface {
	Build(cgin *Cgin)
	Name() string
}
