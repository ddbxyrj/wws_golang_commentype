package wwsTranType

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//给每个模模块的
type WwsConEveryMOdel struct {
	UserId *int
	Status *int
	Uuid   *[]byte
	Log    *log.Logger
}

//每个模块必须实现的函数
type WwsModelPack interface {
	ProcessStringMesg(*[]byte, int8) (int8, *json.RawMessage)
	ProcessByteMesg(*[]byte, int8) (int8, *[]byte)
	SetWwsUserInfo(*WwsConEveryMOdel)
	SetAuth(string) int8
	SetDbGorm(*gorm.DB)
}
