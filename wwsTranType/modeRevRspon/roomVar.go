package modeRevRspon

type RoomDoWhat uint8

const (
	CRoomNew RoomDoWhat = iota + 1
	CRoomMsg
	CRoomJoin
	CRoomJoinAutoCreat
	CRoomLeave
	CRoomDelete
	CRoomSetModeAdmin
)

type Room struct {
	RoomName string
	Dowhat   RoomDoWhat
}
