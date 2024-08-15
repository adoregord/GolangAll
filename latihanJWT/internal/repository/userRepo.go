package repository

import (
	"database/sql"
	"jwt-try/internal/domain"
	utils "jwt-try/internal/utils/hash"
)

type UserRepoInterface interface {
	CheckCredential(user *domain.User) bool
	RegisterUser(user *domain.User) (*domain.User, error)
	CheckVerifiedAcc(username string) (bool, error)
}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepoInterface {
	return UserRepo{
		db: db,
	}
}

func (repo UserRepo) CheckCredential(user *domain.User) bool {
	query := `
	select id, password
	from accounts a
	where username = $1`

	var pass string

	err := repo.db.QueryRow(query, user.Username).Scan(&user.ID, &pass)
	if err != nil {
		return false
	}

	ok := utils.CheckPasswordHash(user.Password, pass)

	return ok
}

func (repo UserRepo) RegisterUser(User *domain.User) (*domain.User, error) {
	query := `
  insert into accounts (username, password)
	values ($1, $2)
  returning id
	`
	err := repo.db.QueryRow(query, User.Username, User.Password).Scan(&User.ID)
	if err != nil {
		return nil, err
	}

	return User, nil
}

func (repo UserRepo) CheckVerifiedAcc(username string) (bool, error) {
	query := `
    select username
    from verified_accounts va 
    where username = $1
  `

	err := repo.db.QueryRow(query, username).Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
