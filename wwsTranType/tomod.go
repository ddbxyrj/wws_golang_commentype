package wwsTranType

type ToModMsgType uint8

const (
	ModSetting = iota + 1
)

type ToModMsg struct {
	MsgType ToModMsgType
}
