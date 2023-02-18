package transactiondto

import (
	"finaltaskbe/models"
)

type CreateTransactionRequest struct {
	Status    string                      `json:"status"`
	AccNumber string                      `json:"acc_Number"`
	OrderDate string                      `json:"order_date"`
	UserID    int                         `json:"user_id" form:"user_id"`
	User      models.UsersProfileResponse `json:"user"`
	FilmID    int                         `json:"film_id" form:"film_id"`
	Film      models.Film                 `json:"film"`
}
