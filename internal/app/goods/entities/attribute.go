package entities

type Attribute struct {
	Id     uint
	Name   string
	Desc   string
	Values []map[uint]string
}
