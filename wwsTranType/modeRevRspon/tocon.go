package modeRevRspon

type ToConDoWhat uint8

const (
	CConClose ToConDoWhat = iota
	CConSentMsgToCon
)

type Tocon struct {
	ConId  []byte
	DoWhat ToConDoWhat
}
