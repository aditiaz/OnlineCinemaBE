package models

type Transaction struct {
	ID        int                  `json:"id"`
	Status    string               `json:"status"`
	AccNumber string               `json:"acc_Number"`
	OrderDate string               `json:"order_date"`
	UserID    int                  `json:"user_id"`
	User      UsersProfileResponse `json:"user" gorm:"foreignKey:UserID"`
	FilmID    int                  `json:"film_id"`
	Film      Film                 `json:"film"`
}

func (Transaction) TableName() string {
	return "transactions"
}
