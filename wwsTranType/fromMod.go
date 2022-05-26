package wwsTranType

type ModAuthBak uint8

const (
	AuthCheckOk ModAuthBak = iota + 1
	AuthVersionLost
	AuthCheckVersionNotSuit
)
