package exporttable

type Export用户表 struct {
	ID        int
	Uid       int `gorm:"Column:user_id"` //UNIQUE
	Name      string
	Show_name string
}

func (Export用户表) TableName() string {
	return "surha_user_true"
}
