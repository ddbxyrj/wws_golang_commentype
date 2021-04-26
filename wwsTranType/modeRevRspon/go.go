package modeRevRspon

type FromeWhere uint8
type RevType uint8
type ToWhere uint8

const (
	FromCon FromeWhere = iota
	From叶集
	FromFace
	FromDharma
	Fromsamsara
)
const (
	RevJson RevType = iota
	RevByte
)

type Recv struct {
	DateByte *[]byte
	From     FromeWhere
	MsgType  RevType
	Dharma   Dharma
	Room     Room
}

const (
	To叶集 ToWhere = iota
	ToDharma
	Tosamsara
	ToConJsonType
	ToConByteType
	ToClose
	ToOtherCon
	NoDo
	ModeError
	ParasError
)

type Resp struct {
	DateByte *[]byte
	To       ToWhere
	MsgType  ToWhere
	Room     Room
	Tocon    Tocon
	Dharma   Dharma
}
