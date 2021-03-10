package wwsTranType

import (
	"encoding/json"
	"fmt"

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
//水流与水面通信模块
type G上达msg struct {
	G连接号 []byte
	G模块好 int
	G内容  []byte
}
