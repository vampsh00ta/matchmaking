package entity

type User struct {
	TgID   int `json:"tg_id" redis:"tg_id"`
	Rating int `json:"rating" redis:"rating"`
}
