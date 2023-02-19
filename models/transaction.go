package models

type Transaction struct {
	ID        int                  `json:"id"`
	Status    string               `json:"status"`
	OrderDate string               `json:"order_date"`
	Price     int                  `json:"price"`
	UserID    int                  `json:"user_id"`
	User      UsersProfileResponse `json:"user" gorm:"foreignKey:UserID"`
	FilmID    int                  `json:"film_id"`
	Film      Film                 `json:"film"`
}

func (Transaction) TableName() string {
	return "transactions"
}
