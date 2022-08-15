package domain

type User struct {
	repository LoginRepository
	Id         int64
	Name       string
	Email      string
	Password   string
	Active     bool
	CreatedAt  string
}

func (user *User) SetRepository(repository LoginRepository) {
	user.repository = repository
}

func (user *User) Exists() bool {
	return user.Id > 0
}

func (user *User) IsActive() bool {
	return user.Active
}

func FindUserByEmail(repository LoginRepository, email string) User {
	return repository.FindUserByEmail(email)
}
