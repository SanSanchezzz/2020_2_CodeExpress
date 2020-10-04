package repositories

import (
	"errors"
	"sync"

	"github.com/go-park-mail-ru/2020_2_CodeExpress/models"
)

type UserRep interface {
	CheckUserExists(u *models.User) error
	CreateUser(u *models.User) error
}

type UserRepImpl struct {
	Users []*models.User
	MU    *sync.RWMutex
}

func NewUserRepImpl() UserRep {
	return &UserRepImpl{
		Users: make([]*models.User, 0),
		MU:    &sync.RWMutex{},
	}
}

func (s *UserRepImpl) CheckUserExists(u *models.User) error {
	s.MU.RLock()
	defer s.MU.RUnlock()
	for _, user := range s.Users {
		if user.Email == u.Email {
			return errors.New("Email already exists")
		}
		if user.Name == u.Name {
			return errors.New("Username already exists")
		}
	}
	return nil
}

func (s *UserRepImpl) CreateUser(u *models.User) error {
	s.MU.Lock()
	defer s.MU.Unlock()
	u.ID = s.getNewUserID()
	s.Users = append(s.Users, u)
	return nil // возвращает nil так как реализация без БД
}

func (s *UserRepImpl) getNewUserID() uint64 {
	if len(s.Users) > 0 {
		return s.Users[len(s.Users)-1].ID+1
	}
	return 0
}