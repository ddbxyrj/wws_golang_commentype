package wwsTranType

import (
	"9.suarha.com/root/wwscommentype.git/wwsTranType/modeRevRspon"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//给每个模模块的
type WwsConEveryMOdel struct {
	UserId *int
	Status *int
	Uuid   *[]byte
	Log    *log.Logger
	Level  UserLevel
}

//每个模块必须实现的函数
type WwsModelPack interface {
	ProcessMesg(*modeRevRspon.Recv) *modeRevRspon.Resp
	// ProcessStringMesg(*[]byte, int8) (int8, *json.RawMessage)
	// ProcessByteMesg(*[]byte, int8) (int8, *[]byte)
	SetWwsUserInfo(*WwsConEveryMOdel)
	SetAuth(*string) ModAuthBak
	SetDbGorm(*gorm.DB)
	SetChan(chan<- modeRevRspon.ModePush, *[16]byte)
	InitModel(s *WwsConEveryMOdel) bool
}
