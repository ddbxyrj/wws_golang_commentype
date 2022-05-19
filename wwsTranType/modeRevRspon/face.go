package modeRevRspon

type WwsDoWahat byte

const (
	InfoConClosed WwsDoWahat = iota + 1
)

func (s *WwsDoWahat) WwsMesgDoWaht(msg *[]byte) (rs *WwsDoWahat) {
	if msg == nil || len(*msg) == 0 {
		var temp WwsDoWahat
		temp = 0
		return &temp
	}
	temp := (WwsDoWahat)((*msg)[0])
	return &temp
}
