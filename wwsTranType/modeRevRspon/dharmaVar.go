package modeRevRspon

type DharmaDoWhat uint8

const (
	DharmaSendMsg DharmaDoWhat = iota
	DharmaInfoSendError
	DharmaINfoRemoteClosed
)

type Dharma struct {
	DharmaId []byte
	Dowhat   DharmaDoWhat
}
