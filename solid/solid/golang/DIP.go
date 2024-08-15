package gosolid

import "github.com/gofrs/uuid"

type UserDataSource interface {
	GetUser(id int) (User, error)
}

type LocalDatabase struct {
}

func (ld *LocalDatabase) GetUser(id int) (User, error) {
	if newUUID, err := uuid.NewV4(); err != nil {
		return User{}, err
	} else {
		uuids := newUUID
		return User{
			sources: "datasource",
			ID:      uuids,
		}, nil
	}
}

type API struct {
}

func (a *API) GetUser(id int) (User, error) {
	if newUUID, err := uuid.NewV4(); err != nil {
		return User{}, err
	} else {
		uuids := newUUID
		return User{
			sources: "apisources",
			ID:      uuids,
		}, nil
	}
}
