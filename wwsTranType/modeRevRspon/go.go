package modeRevRspon

type FromeWhere uint8
type RevType uint8
type ToWhere uint8

const (
	FromCon FromeWhere = iota + 1
	From叶集
	FromFace
	FromDharma
	Fromsamsara
)
const (
	RevJson RevType = iota + 1
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
	To叶集 ToWhere = iota + 1
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

func (s *Resp) ToChan() *RespForChan {
	RespForChan := RespForChan{DateByte: *s.DateByte, To: s.To, MsgType: s.MsgType, Room: s.Room, Tocon: s.Tocon, Dharma: s.Dharma}
	return &RespForChan
}

type RespForChan struct {
	DateByte []byte
	To       ToWhere
	MsgType  ToWhere
	Room     Room
	Tocon    Tocon
	Dharma   Dharma
}
