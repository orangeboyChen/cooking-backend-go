package vo

type UserInfoVO struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Birthday int64  `json:"birthday"`
}
