package model

type UserInfo struct {
	*Model
	UserName string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
	Gender   uint8  `json:"gender"`
}

func (u UserInfo) TableName() string {
	return "blog_user"
}
