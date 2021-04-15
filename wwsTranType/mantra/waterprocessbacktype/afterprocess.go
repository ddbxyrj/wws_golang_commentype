package waterprocessbacktype

type BackMsgTyep int8

const (
	RevJsonBackJson        BackMsgTyep = 0
	RevJsonThenExite       BackMsgTyep = 1
	RevJsonThenToWaterFace BackMsgTyep = 2
	RevJsonThenToRoom      BackMsgTyep = 3
	RevJsonBackByte        BackMsgTyep = 4

	RevByteBackByte BackMsgTyep = 0
)
