package domain

type UserRepositoryInterface interface {
	UserSaver
	UserUpdater
}

type UserSaver interface {
	Save(user *User)
}

type UserUpdater interface {
	Update(user *User)
}

type UserDeleter interface {
	Delete(user *User)
}

// ini domain(model) untuk User
type User struct {
	ID   int
	Name int
	Password string
}
