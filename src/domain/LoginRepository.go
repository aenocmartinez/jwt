package domain

type LoginRepository interface {
	FindUserByEmail(email string) User
}
