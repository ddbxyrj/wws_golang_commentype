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
	// SetAuth(*[]byte)
	// authToString() string
	ProcessMesg(*[]byte) (int, *json.RawMessage)
	SetWwsUserInfo(*WwsConEveryMOdel)
	SetAuth(string) int
	SetDbGorm(*gorm.DB)
}

