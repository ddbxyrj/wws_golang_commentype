package modeRevRspon

type WwsDoWahat struct {
	RevDate *[]byte
}

type WwsDo byte

const (
	InfoConClosed WwsDo = iota + 1
)

func (s *WwsDoWahat) WwsMesgDoWaht() (rs *WwsDo) {
	if s.RevDate == nil || len(*s.RevDate) == 0 {
		var temp WwsDo
		temp = 0
		return &temp
	}
	temp := (WwsDo)((*s.RevDate)[0])
	return &temp
}
