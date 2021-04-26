package modeRevRspon

type ToConDoWhat uint8

const (
	CConClose ToConDoWhat = iota + 1
	CConSentMsgToCon
	CGetAllCon
)

type Tocon struct {
	ConId  []byte
	DoWhat ToConDoWhat
}
