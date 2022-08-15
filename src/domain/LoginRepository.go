package domain

type LoginRepository interface {
	FindUserByEmail(email string) User
	UpdateToken(idUser int64, token string) bool
}
