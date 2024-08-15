package usecase

import (
	"basic-solid/internal/domain"
	"basic-solid/internal/repository"
	"errors"
)

type UserUsecase struct {
	userRepo    repository.UserRepository
	oauthServer repository.UserOauthServer
}

// TODO: This need move to package dto if possible? cmmiw? yesss 
type LoginRequest struct {
	UserId int `json:"user_id"`
	Password string `json:"password"`
	isOauth  bool   `json:"-"` // this should ignore in json, user cannot send
}

type UserUsecaseInterface interface {
	Login(request LoginRequest) error
}

type UserUsecaseImpl struct {
	userRepo repository.UserDatasource
}

func NewUserUsecase(
	userRepo repository.UserRepository,
	userOauth repository.UserOauthServer,
) UserUsecaseInterface {
	return &UserUsecaseImpl{
		userRepo:    userRepo,
		oauthServer: userOauth,
	}
}

// ini fungsi untuk login 😍🤣😂😊❤️💕🙌👌🤦‍♀️
func (u *UserUsecaseImpl) Login(request LoginRequest) error {
	var user *domain.User
	var err error
	// if the isOauth 'true', should get data user from oauth server
	// and if its not, do otherwise get from db
	if request.isOauth {
		user, err = u.oauthServer.Find(request.UserId)

		if err != nil {
			return err
		}
	}  

	user, err = u.userRepo.Find(request.UserId)

	// if the password is not match, return an error 
	if request.Password != user.Password {
		return errors.New("Invalid password")
	} 

	return nil
}

