package models

type User struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"  gorm:"type:varchar(255);unique;not null"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
type UsersProfileResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
