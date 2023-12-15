package models

type Feed struct {
	UserId int64  `json:"user_id, omitempty"`
	Title  string `json:"title, omitempty"`
	Desc   string `json:"desc, omitempty"`
}
